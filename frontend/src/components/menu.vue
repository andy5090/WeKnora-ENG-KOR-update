<template>
    <div class="aside_box">
        <div class="logo_box" @click="router.push('/platform/knowledge-bases')" style="cursor: pointer;">
            <img class="logo" src="@/assets/img/weknora.png" alt="">
        </div>
        
        <!-- Tenant selector: Only shown when user can switch tenants -->
        <TenantSelector v-if="canAccessAllTenants" />
        
        <!-- Top section: Knowledge bases and chats -->
        <div class="menu_top">
            <div class="menu_box" :class="{ 'has-submenu': item.children }" v-for="(item, index) in topMenuItems" :key="index">
                <div @click="handleMenuClick(item.path)"
                    @mouseenter="mouseenteMenu(item.path)" @mouseleave="mouseleaveMenu(item.path)"
                     :class="['menu_item', item.childrenPath && item.childrenPath == currentpath ? 'menu_item_c_active' : isMenuItemActive(item.path) ? 'menu_item_active' : '']">
                    <div class="menu_item-box">
                        <div class="menu_icon">
                            <img class="icon" :src="getImgSrc(item.icon == 'zhishiku' ? knowledgeIcon : item.icon == 'agent' ? agentIcon : item.icon == 'organization' ? organizationIcon : item.icon == 'logout' ? logoutIcon : item.icon == 'setting' ? settingIcon : prefixIcon)" alt="">
                        </div>
                        <span class="menu_title" :title="item.title">{{ item.title }}</span>
                        <span v-if="item.path === 'organizations' && orgStore.totalPendingJoinRequestCount > 0" class="menu-pending-badge" :title="t('organization.settings.pendingJoinRequestsBadge')">{{ orgStore.totalPendingJoinRequestCount }}</span>
                        <t-icon v-if="item.path === 'creatChat'" name="add" class="menu-create-hint" />
                    </div>
                </div>
                <div ref="submenuscrollContainer" @scroll="handleScroll" class="submenu" v-if="item.children">
                    <template v-for="(group, groupIndex) in groupedSessions" :key="groupIndex">
                        <div class="timeline_header">{{ group.label }}</div>
                        <div class="submenu_item_p" v-for="(subitem, subindex) in group.items" :key="subitem.id">
                            <div :class="['submenu_item', currentSecondpath == subitem.path ? 'submenu_item_active' : '']"
                                @mouseenter="mouseenteBotDownr(subitem.id)" @mouseleave="mouseleaveBotDown"
                                @click="gotopage(subitem.path)">
                                <span class="submenu_title"
                                    :style="currentSecondpath == subitem.path ? 'margin-left:18px;max-width:160px;' : 'margin-left:18px;max-width:185px;'">
                                    {{ subitem.title }}
                                </span>
                                <t-dropdown 
                                    :options="[{ content: t('upload.deleteRecord'), value: 'delete' }, { content: t('menu.batchManage'), value: 'batchManage' }]"
                                    @click="handleSessionMenuClick($event, subitem.originalIndex, subitem)"
                                    placement="bottom-right"
                                    trigger="click">
                                    <div @click.stop class="menu-more-wrap">
                                        <t-icon name="ellipsis" class="menu-more" />
                                    </div>
                                </t-dropdown>
                            </div>
                        </div>
                    </template>
                </div>
            </div>
        </div>
        
        
        <!-- Bottom section: User menu -->
        <div class="menu_bottom">
            <UserMenu />
        </div>

        <!-- 批量管理对话框 -->
        <BatchManageDialog
            v-model:visible="batchManageVisible"
            :sessions="allSessions"
            @deleted="handleBatchDeleted"
        />
        
    </div>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia';
import { onMounted, watch, computed, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getSessionsList, delSession } from "@/api/chat/index";
import { getKnowledgeBaseById } from '@/api/knowledge-base';
import { logout as logoutApi } from '@/api/auth';
import { useMenuStore } from '@/stores/menu';
import { useAuthStore } from '@/stores/auth';
import { useOrganizationStore } from '@/stores/organization';
import { useUIStore } from '@/stores/ui';
import { MessagePlugin } from "tdesign-vue-next";
import UserMenu from '@/components/UserMenu.vue';
import TenantSelector from '@/components/TenantSelector.vue';
import BatchManageDialog from '@/components/BatchManageDialog.vue';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const usemenuStore = useMenuStore();
const authStore = useAuthStore();
const orgStore = useOrganizationStore();
const uiStore = useUIStore();
const route = useRoute();
const router = useRouter();
const currentpath = ref('');
const currentPage = ref(1);
const page_size = ref(30);
const total = ref(0);
const currentSecondpath = ref('');
const submenuscrollContainer = ref(null);
// Calculate total pages
const totalPages = computed(() => Math.ceil(total.value / page_size.value));
const hasMore = computed(() => currentPage.value < totalPages.value);
type MenuItem = { title: string; icon: string; path: string; childrenPath?: string; children?: any[] };
const { menuArr } = storeToRefs(usemenuStore);
let activeSubmenu = ref<string>('');

// Batch management state
const batchManageVisible = ref(false);

// All sessions list (for batch management dialog)
const allSessions = computed(() => {
    const chatMenu = (menuArr.value as unknown as MenuItem[]).find((item: MenuItem) => item.path === 'creatChat');
    if (!chatMenu || !chatMenu.children) return [];
    return chatMenu.children.map((s: any) => ({
        id: s.id,
        title: s.title,
        created_at: s.created_at,
        updated_at: s.updated_at,
    }));
});

// Whether user can access all tenants
const canAccessAllTenants = computed(() => authStore.canAccessAllTenants);

// Whether on knowledge base detail page (excluding global chat)
const isInKnowledgeBase = computed<boolean>(() => {
    return route.name === 'knowledgeBaseDetail' || 
           route.name === 'kbCreatChat' || 
           route.name === 'knowledgeBaseSettings';
});

// Whether on knowledge base list page
const isInKnowledgeBaseList = computed<boolean>(() => {
    return route.name === 'knowledgeBaseList';
});

// Whether on create chat page
const isInCreatChat = computed<boolean>(() => {
    return route.name === 'globalCreatChat' || route.name === 'kbCreatChat';
});

// Whether on chat detail page
const isInChatDetail = computed<boolean>(() => route.name === 'chat');

// Whether on agent list page
const isInAgentList = computed<boolean>(() => route.name === 'agentList');

// Whether on organization list page
const isInOrganizationList = computed<boolean>(() => route.name === 'organizationList');

// Unified menu item active state determination
const isMenuItemActive = (itemPath: string): boolean => {
    const currentRoute = route.name;
    
    switch (itemPath) {
        case 'knowledge-bases':
            return currentRoute === 'knowledgeBaseList' || 
                   currentRoute === 'knowledgeBaseDetail' || 
                   currentRoute === 'knowledgeBaseSettings';
        case 'agents':
            return currentRoute === 'agentList';
        case 'organizations':
            return currentRoute === 'organizationList';
        case 'creatChat':
            return currentRoute === 'kbCreatChat' || currentRoute === 'globalCreatChat';
        case 'settings':
            return currentRoute === 'settings';
        default:
            return itemPath === currentpath.value;
    }
};

// Unified icon active state determination
const getIconActiveState = (itemPath: string) => {
    const currentRoute = route.name;
    
    return {
        isKbActive: itemPath === 'knowledge-bases' && (
            currentRoute === 'knowledgeBaseList' || 
            currentRoute === 'knowledgeBaseDetail' || 
            currentRoute === 'knowledgeBaseSettings'
        ),
        isCreatChatActive: itemPath === 'creatChat' && (currentRoute === 'kbCreatChat' || currentRoute === 'globalCreatChat'),
        isSettingsActive: itemPath === 'settings' && currentRoute === 'settings',
        isChatActive: itemPath === 'chat' && currentRoute === 'chat'
    };
};

// Separate top and bottom menu sections
const topMenuItems = computed<MenuItem[]>(() => {
    return (menuArr.value as unknown as MenuItem[]).filter((item: MenuItem) => 
        item.path === 'knowledge-bases' || item.path === 'agents' || item.path === 'organizations' || item.path === 'creatChat'
    );
});

const bottomMenuItems = computed<MenuItem[]>(() => {
    return (menuArr.value as unknown as MenuItem[]).filter((item: MenuItem) => {
        if (item.path === 'knowledge-bases' || item.path === 'agents' || item.path === 'organizations' || item.path === 'creatChat') {
            return false;
        }
        return true;
    });
});

// Current knowledge base information
const currentKbName = ref<string>('')
const currentKbInfo = ref<any>(null)
const docUploadInput = ref<HTMLInputElement | null>(null)
const docFolderInput = ref<HTMLInputElement | null>(null)
const pendingUploadKbId = ref<string | null>(null)
const selectedFaqCount = ref<number>(0)
const selectedFaqEnabledCount = ref<number>(0)
const selectedFaqDisabledCount = ref<number>(0)

// Listen for FAQ selection count changes
const handleFaqSelectionChanged = ((event: CustomEvent<{ count: number; enabledCount?: number; disabledCount?: number }>) => {
  const count = event.detail?.count || 0
  selectedFaqCount.value = count
  selectedFaqEnabledCount.value = event.detail?.enabledCount || 0
  selectedFaqDisabledCount.value = event.detail?.disabledCount || 0
}) as EventListener

const showKbActions = computed(() => 
    (isInKnowledgeBase.value && !!currentKbInfo.value) || 
    isInKnowledgeBaseList.value || 
    isInCreatChat.value ||
    isInChatDetail.value ||
    isInAgentList.value
)
const currentKbType = computed(() => currentKbInfo.value?.type || 'document')
const showDocActions = computed(() => showKbActions.value && isInKnowledgeBase.value && currentKbType.value !== 'faq')
const showFaqActions = computed(() => showKbActions.value && isInKnowledgeBase.value && currentKbType.value === 'faq')
const showCreateKbAction = computed(() => showKbActions.value && (isInKnowledgeBaseList.value || isInCreatChat.value || isInChatDetail.value))
const showCreateAgentAction = computed(() => showKbActions.value && isInAgentList.value)

// Time grouping function
const getTimeCategory = (dateStr: string): string => {
    if (!dateStr) return t('time.earlier');
    
    const date = new Date(dateStr);
    const now = new Date();
    const today = new Date(now.getFullYear(), now.getMonth(), now.getDate());
    const yesterday = new Date(today.getTime() - 24 * 60 * 60 * 1000);
    const sevenDaysAgo = new Date(today.getTime() - 7 * 24 * 60 * 60 * 1000);
    const thirtyDaysAgo = new Date(today.getTime() - 30 * 24 * 60 * 60 * 1000);
    const oneYearAgo = new Date(today.getTime() - 365 * 24 * 60 * 60 * 1000);
    
    const sessionDate = new Date(date.getFullYear(), date.getMonth(), date.getDate());
    
    if (sessionDate.getTime() >= today.getTime()) {
        return t('time.today');
    } else if (sessionDate.getTime() >= yesterday.getTime()) {
        return t('time.yesterday');
    } else if (date.getTime() >= sevenDaysAgo.getTime()) {
        return t('time.last7Days');
    } else if (date.getTime() >= thirtyDaysAgo.getTime()) {
        return t('time.last30Days');
    } else if (date.getTime() >= oneYearAgo.getTime()) {
        return t('time.lastYear');
    } else {
        return t('time.earlier');
    }
};

// Group session list by time
const groupedSessions = computed(() => {
    const chatMenu = (menuArr.value as unknown as MenuItem[]).find((item: MenuItem) => item.path === 'creatChat');
    if (!chatMenu || !chatMenu.children || chatMenu.children.length === 0) {
        return [];
    }
    
    const groups: { [key: string]: any[] } = {
        [t('time.today')]: [],
        [t('time.yesterday')]: [],
        [t('time.last7Days')]: [],
        [t('time.last30Days')]: [],
        [t('time.lastYear')]: [],
        [t('time.earlier')]: []
    };
    
    // Group sessions by time
    (chatMenu.children as any[]).forEach((session: any, index: number) => {
        const category = getTimeCategory(session.updated_at || session.created_at);
        groups[category].push({
            ...session,
            originalIndex: index
        });
    });
    
    // Return non-empty groups in order
    const orderedLabels = [t('time.today'), t('time.yesterday'), t('time.last7Days'), t('time.last30Days'), t('time.lastYear'), t('time.earlier')];
    return orderedLabels
        .filter(label => groups[label].length > 0)
        .map(label => ({
            label,
            items: groups[label]
        }));
});

const loading = ref(false)
const mouseenteBotDownr = (val: string) => {
    activeSubmenu.value = val;
}
const mouseleaveBotDown = () => {
    activeSubmenu.value = '';
}

const handleSessionMenuClick = (data: { value: string }, index: number, item: any) => {
    if (data?.value === 'delete') {
        delCard(index, item);
    } else if (data?.value === 'batchManage') {
        batchManageVisible.value = true;
    }
};

const delCard = (index: number, item: any) => {
    delSession(item.id).then((res: any) => {
        if (res && (res as any).success) {
            // Find 'creatChat' menu item
            const chatMenuItem = (menuArr.value as any[]).find((m: any) => m.path === 'creatChat');
            
            if (chatMenuItem && chatMenuItem.children) {
                const children = chatMenuItem.children;
                // Find index by ID, safer than relying on passed index
                const actualIndex = children.findIndex((s: any) => s.id === item.id);
                
                if (actualIndex !== -1) {
                    children.splice(actualIndex, 1);
                }
            }
            
            if (item.id == route.params.chatid) {
                // After deleting current session, navigate to global create chat page
                router.push('/platform/creatChat');
            }
            // Update total count
            if (total.value > 0) {
                total.value--;
            }
        } else {
            MessagePlugin.error("Delete failed, please try again later!");
        }
    })
}

const handleBatchDeleted = (ids: string[]) => {
    const chatMenuItem = (menuArr.value as any[]).find((m: any) => m.path === 'creatChat');
    if (chatMenuItem && chatMenuItem.children) {
        const children = chatMenuItem.children;
        for (const id of ids) {
            const idx = children.findIndex((s: any) => s.id === id);
            if (idx !== -1) children.splice(idx, 1);
        }
    }
    total.value = Math.max(0, total.value - ids.length);
    // 如果当前会话被删除，跳转到创建页
    const currentChatId = route.params.chatid as string;
    if (currentChatId && ids.includes(currentChatId)) {
        router.push('/platform/creatChat');
    }
}

const debounce = (fn: (...args: any[]) => void, delay: number) => {
    let timer: ReturnType<typeof setTimeout>
    return (...args: any[]) => {
        clearTimeout(timer)
        timer = setTimeout(() => fn(...args), delay)
    }
}
// Scroll handling
const checkScrollBottom = () => {
    const container = submenuscrollContainer.value
    if (!container || !container[0]) return

    const { scrollTop, scrollHeight, clientHeight } = container[0]
    const isBottom = scrollHeight - (scrollTop + clientHeight) < 100 // Bottom threshold
    
    if (isBottom && hasMore.value && !loading.value) {
        currentPage.value++;
        getMessageList(true);
    }
}
const handleScroll = debounce(checkScrollBottom, 200)
const getMessageList = async (isLoadMore = false) => {
    if (loading.value) return Promise.resolve();
    loading.value = true;
    
    // Only clear array on first load or route change, don't clear on scroll load
    if (!isLoadMore) {
        currentPage.value = 1; // Reset page number
        usemenuStore.clearMenuArr();
    }
    
    return getSessionsList(currentPage.value, page_size.value).then((res: any) => {
        if (res.data && res.data.length) {
            // Display all sessions globally without filtering
            res.data.forEach((item: any) => {
                let obj = { 
                    title: item.title ? item.title : "New Session", 
                    path: `chat/${item.id}`, 
                    id: item.id, 
                    isMore: false, 
                    isNoTitle: item.title ? false : true,
                    created_at: item.created_at,
                    updated_at: item.updated_at
                }
                usemenuStore.updatemenuArr(obj)
            });
        }
        if ((res as any).total) {
            total.value = (res as any).total;
        }
        loading.value = false;
    }).catch(() => {
        loading.value = false;
    })
}

onMounted(async () => {
    const routeName = typeof route.name === 'string' ? route.name : (route.name ? String(route.name) : '')
    currentpath.value = routeName;
    if (route.params.chatid) {
        currentSecondpath.value = `chat/${route.params.chatid}`;
    }
    
    // Initialize knowledge base information
    const kbId = (route.params as any)?.kbId as string
    if (kbId && isInKnowledgeBase.value) {
        try {
            const kbRes: any = await getKnowledgeBaseById(kbId)
            if (kbRes?.data) {
                currentKbName.value = kbRes.data.name || ''
                currentKbInfo.value = kbRes.data
            }
        } catch {}
    } else {
        currentKbName.value = ''
        currentKbInfo.value = null
    }
    
    // Load chat list
    getMessageList();
    // If organization list not loaded, fetch once for sidebar "pending approval" badge
    if (orgStore.organizations.length === 0) {
        orgStore.fetchOrganizations();
    }
    // Listen for FAQ selection count changes
    window.addEventListener('faqSelectionChanged', handleFaqSelectionChanged)
});

watch([() => route.name, () => route.params], (newvalue, oldvalue) => {
    // Reset selection count when switching knowledge base
    if (newvalue[1].kbId !== oldvalue?.[1]?.kbId) {
        selectedFaqCount.value = 0
    }
    const nameStr = typeof newvalue[0] === 'string' ? (newvalue[0] as string) : (newvalue[0] ? String(newvalue[0]) : '')
    currentpath.value = nameStr;
    if (newvalue[1].chatid) {
        currentSecondpath.value = `chat/${newvalue[1].chatid}`;
    } else {
        currentSecondpath.value = "";
    }
    
    // Only refresh chat list when necessary to avoid unnecessary reloads causing list jitter
    // Cases requiring refresh:
    // 1. After creating new session (navigating from creatChat/kbCreatChat to chat/:id)
    // 2. After deleting session, already handled in delCard, no need to refresh here
    const oldRouteNameStr = typeof oldvalue?.[0] === 'string' ? (oldvalue[0] as string) : (oldvalue?.[0] ? String(oldvalue[0]) : '')
    const isCreatingNewSession = (oldRouteNameStr === 'globalCreatChat' || oldRouteNameStr === 'kbCreatChat') && 
                                 nameStr !== 'globalCreatChat' && nameStr !== 'kbCreatChat';
    
    // Only refresh list when creating new session
    if (isCreatingNewSession) {
        getMessageList();
    }
    
    // Update icon state and knowledge base info on route change (not involving chat list)
    getIcon(nameStr);
    
    // If knowledge base switched, update name but don't reload chat list
    if (newvalue[1].kbId !== oldvalue?.[1]?.kbId) {
        const kbId = (newvalue[1] as any)?.kbId as string;
        if (kbId && isInKnowledgeBase.value) {
            getKnowledgeBaseById(kbId).then((kbRes: any) => {
                if (kbRes?.data) {
                    currentKbName.value = kbRes.data.name || '';
                    currentKbInfo.value = kbRes.data;
                }
            }).catch(() => {
                currentKbInfo.value = null;
            });
        } else {
            currentKbName.value = '';
            currentKbInfo.value = null;
        }
    }
});
let knowledgeIcon = ref('zhishiku-green.svg');
let prefixIcon = ref('prefixIcon.svg');
let logoutIcon = ref('logout.svg');
let settingIcon = ref('setting.svg'); // Settings icon
let agentIcon = ref('agent.svg'); // Agent icon
let organizationIcon = ref('organization.svg'); // Organization icon
let pathPrefix = ref(route.name)
  const getIcon = (path: string) => {
      // Update all icons based on current route state
      const kbActiveState = getIconActiveState('knowledge-bases');
      const creatChatActiveState = getIconActiveState('creatChat');
      const settingsActiveState = getIconActiveState('settings');
      const agentsActiveState = route.name === 'agentList';
      const organizationsActiveState = route.name === 'organizationList';
      
      // Knowledge base icon: Only show green on knowledge base pages
      knowledgeIcon.value = kbActiveState.isKbActive ? 'zhishiku-green.svg' : 'zhishiku.svg';
      
      // Agent icon: Only show green on agent pages
      agentIcon.value = agentsActiveState ? 'agent-green.svg' : 'agent.svg';
      
      // Organization icon: Only show green on organization pages
      organizationIcon.value = organizationsActiveState ? 'organization-green.svg' : 'organization.svg';
      
      // Chat icon: Only show green on chat creation pages, grey on knowledge base pages, default otherwise
      prefixIcon.value = creatChatActiveState.isCreatChatActive ? 'prefixIcon-green.svg' : 
                        kbActiveState.isKbActive ? 'prefixIcon-grey.svg' : 
                        'prefixIcon.svg';
      
      // Settings icon: Only show green on settings page
      settingIcon.value = settingsActiveState.isSettingsActive ? 'setting-green.svg' : 'setting.svg';
      
      // Logout icon: Always show default
      logoutIcon.value = 'logout.svg';
}
getIcon(typeof route.name === 'string' ? route.name as string : (route.name ? String(route.name) : ''))
const handleMenuClick = async (path: string) => {
    if (path === 'knowledge-bases') {
        // Knowledge base menu item: If inside knowledge base, navigate to current knowledge base files page; otherwise navigate to knowledge base list
        const kbId = await getCurrentKbId()
        if (kbId) {
            router.push(`/platform/knowledge-bases/${kbId}`)
        } else {
            router.push('/platform/knowledge-bases')
        }
    } else if (path === 'agents') {
        // Agent menu item: Navigate to agent list
        router.push('/platform/agents')
    } else if (path === 'organizations') {
        // Organization menu item: Navigate to organization list
        router.push('/platform/organizations')
    } else if (path === 'settings') {
        // Settings menu item: Open settings modal and navigate route
        uiStore.openSettings()
        router.push('/platform/settings')
    } else {
        gotopage(path)
    }
}

// Handle logout confirmation
const handleLogout = () => {
    gotopage('logout')
}

const getCurrentKbId = async (): Promise<string | null> => {
    const kbId = (route.params as any)?.kbId as string
    if (isInKnowledgeBase.value && kbId) {
        return kbId
    }
    return null
}

const gotopage = async (path: string) => {
    pathPrefix.value = path;
    // Handle logout
    if (path === 'logout') {
        try {
            // Call backend API to logout
            await logoutApi();
        } catch (error) {
            // Even if API call fails, continue with local cleanup
            console.error('Logout API call failed:', error);
        }
        // Clear all state and local storage
        authStore.logout();
        MessagePlugin.success('Logged out successfully');
        router.push('/login');
        return;
    } else {
        if (path === 'creatChat') {
            // If on knowledge base detail page, navigate to global chat creation page
            if (isInKnowledgeBase.value) {
                router.push('/platform/creatChat')
            } else {
                // If not inside knowledge base, enter chat creation page
                router.push(`/platform/creatChat`)
            }
        } else {
            router.push(`/platform/${path}`);
        }
    }
    getIcon(path)
}

const getImgSrc = (url: string) => {
    return new URL(`/src/assets/img/${url}`, import.meta.url).href;
}

const mouseenteMenu = (path: string) => {
    if (pathPrefix.value != 'knowledge-bases' && pathPrefix.value != 'creatChat' && path != 'knowledge-bases') {
        prefixIcon.value = 'prefixIcon-grey.svg';
    }
}
const mouseleaveMenu = (path: string) => {
    if (pathPrefix.value != 'knowledge-bases' && pathPrefix.value != 'creatChat' && path != 'knowledge-bases') {
        const nameStr = typeof route.name === 'string' ? route.name as string : (route.name ? String(route.name) : '')
        getIcon(nameStr)
    }
}

const ensureDocKnowledgeBaseReady = async (): Promise<string | null> => {
    const kbId = await getCurrentKbId()
    if (!kbId) {
        MessagePlugin.warning(t('knowledgeEditor.messages.missingId'))
        return null
    }
    if (currentKbType.value === 'faq') {
        MessagePlugin.warning(t('knowledgeBase.docActionUnsupported'))
        return null
    }
    if (!currentKbInfo.value || !currentKbInfo.value.embedding_model_id || !currentKbInfo.value.summary_model_id) {
        MessagePlugin.warning(t('knowledgeBase.notInitialized'))
        return null
    }
    return kbId
}

const handleDocUploadClick = async () => {
    const kbId = await ensureDocKnowledgeBaseReady()
    if (!kbId) return
    pendingUploadKbId.value = kbId
    docUploadInput.value?.click()
}

const FAILED_FILES_PREVIEW_LIMIT = 10

const summarizeFailedFiles = (failedFiles: Array<{ name: string; reason: string }>) => {
    const duplicateLabel = t('knowledgeBase.fileExists')
    let duplicateCount = 0
    const nonDuplicate: Array<{ name: string; reason: string }> = []
    failedFiles.forEach((file) => {
        if (file.reason === duplicateLabel) {
            duplicateCount++
        } else {
            nonDuplicate.push(file)
        }
    })

    const previewList = nonDuplicate.slice(0, FAILED_FILES_PREVIEW_LIMIT).map(f => `• ${f.name}: ${f.reason}`)
    let nonDuplicateText = ''
    if (previewList.length) {
        nonDuplicateText = previewList.join('\n')
        if (nonDuplicate.length > FAILED_FILES_PREVIEW_LIMIT) {
            nonDuplicateText += `\n${t('knowledgeBase.andMoreFiles', { count: nonDuplicate.length - FAILED_FILES_PREVIEW_LIMIT })}`
        }
    }

    return {
        duplicateCount,
        nonDuplicateText,
    }
}

const handleDocFileChange = async (event: Event) => {
    const input = event.target as HTMLInputElement
    const files = input?.files
    if (!files || files.length === 0) {
        pendingUploadKbId.value = null
        return
    }

    const kbId = pendingUploadKbId.value || (await ensureDocKnowledgeBaseReady())
    pendingUploadKbId.value = null
    if (!kbId) {
        input.value = ''
        return
    }

    // Filter valid files
    const validFiles: File[] = []
    let invalidCount = 0
    const isSingleFile = files.length === 1

    for (let i = 0; i < files.length; i++) {
        const file = files[i]
        // Show error for single file, silently filter for multiple files
        if (kbFileTypeVerification(file, !isSingleFile)) {
            invalidCount++
        } else {
            validFiles.push(file)
        }
    }

    // If no valid files, show summary prompt for multiple files
    if (validFiles.length === 0) {
        if (!isSingleFile && invalidCount > 0) {
            MessagePlugin.error(t('knowledgeBase.noValidFilesSelected'))
        }
        // Single file errors are already shown in kbFileTypeVerification
        input.value = ''
        return
    }

    // Batch upload
    let successCount = 0
    let failCount = 0
    const totalCount = validFiles.length
    const failedFiles: Array<{ name: string; reason: string }> = []

    // Create upload task for each file and send event notification
    const uploadPromises = validFiles.map(async (file) => {
        const uploadId = `${file.name}_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`
        let progress = 0
        let status: 'uploading' | 'success' | 'error' = 'uploading'
        let error: string | undefined

        // Send upload start event
        window.dispatchEvent(new CustomEvent('knowledgeFileUploadStart', {
            detail: { 
                kbId, 
                uploadId, 
                fileName: file.name
            }
        }))

        try {
            // Get currently selected tag ID
            const tagIdToUpload = uiStore.selectedTagId !== '__untagged__' ? uiStore.selectedTagId : undefined
            await uploadKnowledgeFile(
                kbId, 
                { file, tag_id: tagIdToUpload },
                (progressEvent: any) => {
                    if (progressEvent.total) {
                        progress = Math.round((progressEvent.loaded * 100) / progressEvent.total)
                        // Send progress update event
                        window.dispatchEvent(new CustomEvent('knowledgeFileUploadProgress', {
                            detail: { 
                                kbId, 
                                uploadId, 
                                progress 
                            }
                        }))
                    }
                }
            )
            successCount++
            status = 'success'
            progress = 100
        } catch (error: any) {
            failCount++
            let errorReason = error?.error?.message || error?.message || t('knowledgeBase.uploadFailed')
            if (error?.code === 'duplicate_file' || error?.error?.code === 'duplicate_file') {
                errorReason = t('knowledgeBase.fileExists')
            }
            status = 'error'
            error = errorReason
            failedFiles.push({ name: file.name, reason: errorReason })

            // Only show detailed error for single file upload
            if (totalCount === 1) {
                MessagePlugin.error(errorReason)
            }
        } finally {
            // Send upload complete event
            window.dispatchEvent(new CustomEvent('knowledgeFileUploadComplete', {
                detail: { 
                    kbId, 
                    uploadId, 
                    status,
                    progress,
                    error
                }
            }))
        }
    })

    // Wait for all uploads to complete
    await Promise.allSettled(uploadPromises)

    // Show upload results
    if (successCount > 0) {
        window.dispatchEvent(new CustomEvent('knowledgeFileUploaded', {
            detail: { kbId }
        }))
    }

    if (totalCount === 1) {
        if (successCount === 1) {
            MessagePlugin.success(t('knowledgeBase.uploadSuccess'))
        }
        // Single file failure errors are already shown above
    } else {
        if (failCount === 0) {
            MessagePlugin.success(t('knowledgeBase.uploadAllSuccess', { count: successCount }))
        } else if (successCount > 0) {
            const { duplicateCount, nonDuplicateText } = summarizeFailedFiles(failedFiles)
            const extraSections: string[] = []
            if (duplicateCount > 0) {
                extraSections.push(t('knowledgeBase.duplicateFilesSkipped', { count: duplicateCount }))
            }
            if (nonDuplicateText) {
                extraSections.push(t('knowledgeBase.failedFilesList') + '\n' + nonDuplicateText)
            }
            const extraContent = extraSections.length ? '\n\n' + extraSections.join('\n\n') : ''
            MessagePlugin.warning({
                content: t('knowledgeBase.uploadPartialSuccess', {
                    success: successCount,
                    fail: failCount
                }) + extraContent,
                duration: 8000,
                closeBtn: true
            })
        } else {
            const { duplicateCount, nonDuplicateText } = summarizeFailedFiles(failedFiles)
            const extraSections: string[] = []
            if (duplicateCount > 0) {
                extraSections.push(t('knowledgeBase.duplicateFilesSkipped', { count: duplicateCount }))
            }
            if (nonDuplicateText) {
                extraSections.push(t('knowledgeBase.failedFilesList') + '\n' + nonDuplicateText)
            }
            const extraContent = extraSections.length ? '\n\n' + extraSections.join('\n\n') : ''
            MessagePlugin.error({
                content: t('knowledgeBase.uploadAllFailed') + extraContent,
                duration: 8000,
                closeBtn: true
            })
        }
    }

    input.value = ''
}

const handleDocFolderUploadClick = async () => {
    const kbId = await ensureDocKnowledgeBaseReady()
    if (!kbId) return
    pendingUploadKbId.value = kbId
    docFolderInput.value?.click()
}

const handleDocFolderChange = async (event: Event) => {
    const input = event.target as HTMLInputElement
    const files = input?.files
    if (!files || files.length === 0) {
        pendingUploadKbId.value = null
        return
    }

    const kbId = pendingUploadKbId.value || (await ensureDocKnowledgeBaseReady())
    pendingUploadKbId.value = null
    if (!kbId) {
        input.value = ''
        return
    }

    // Check if VLM is enabled
    const vlmEnabled = currentKbInfo.value?.vlm_config?.enabled || false

    // Filter valid files (folder upload always uses silent mode)
    const validFiles: File[] = []
    let invalidCount = 0
    let hiddenFileCount = 0
    let imageFilteredCount = 0

    for (let i = 0; i < files.length; i++) {
        const file = files[i]
        const relativePath = (file as any).webkitRelativePath || file.name
        
        // 1. Filter hidden files and hidden folders
        // Check if path contains files or folders starting with .
        const pathParts = relativePath.split('/')
        const hasHiddenComponent = pathParts.some((part: string) => part.startsWith('.'))
        if (hasHiddenComponent) {
            hiddenFileCount++
            continue
        }
        
        // 2. If VLM not enabled, filter image files
        if (!vlmEnabled) {
            const fileExt = file.name.substring(file.name.lastIndexOf('.') + 1).toLowerCase()
            const imageTypes = ['jpg', 'jpeg', 'png', 'gif', 'bmp', 'webp']
            if (imageTypes.includes(fileExt)) {
                imageFilteredCount++
                continue
            }
        }
        
        // 3. File type validation (always silently filter for folder upload)
        if (kbFileTypeVerification(file, true)) {
            invalidCount++
        } else {
            validFiles.push(file)
        }
    }

    // If no valid files, return directly
    if (validFiles.length === 0) {
        const totalFiltered = invalidCount + hiddenFileCount + imageFilteredCount
        if (totalFiltered > 0) {
            let filterReasons = []
            if (hiddenFileCount > 0) {
                filterReasons.push(t('knowledgeBase.hiddenFilesFiltered', { count: hiddenFileCount }))
            }
            if (imageFilteredCount > 0) {
                filterReasons.push(t('knowledgeBase.imagesFilteredNoVLM', { count: imageFilteredCount }))
            }
            if (invalidCount > 0) {
                filterReasons.push(t('knowledgeBase.invalidFilesFiltered', { count: invalidCount }))
            }
            MessagePlugin.warning(t('knowledgeBase.noValidFilesInFolder', { total: files.length }) + '\n' + filterReasons.join('\n'))
        } else {
            MessagePlugin.error(t('knowledgeBase.noValidFiles'))
        }
        input.value = ''
        return
    }

    // Show filtered upload prompt
    const totalCount = validFiles.length
    const totalFiltered = invalidCount + hiddenFileCount + imageFilteredCount
    if (totalFiltered > 0) {
        let filterInfo = []
        if (hiddenFileCount > 0) {
            filterInfo.push(t('knowledgeBase.hiddenFilesFiltered', { count: hiddenFileCount }))
        }
        if (imageFilteredCount > 0) {
            filterInfo.push(t('knowledgeBase.imagesFilteredNoVLM', { count: imageFilteredCount }))
        }
        if (invalidCount > 0) {
            filterInfo.push(t('knowledgeBase.invalidFilesFiltered', { count: invalidCount }))
        }
        MessagePlugin.info(
            t('knowledgeBase.uploadingValidFiles', {
                valid: totalCount,
                total: files.length
            }) + '\n' + filterInfo.join(', ')
        )
    } else {
        MessagePlugin.info(t('knowledgeBase.uploadingFolder', { total: totalCount }))
    }

    // Batch upload folder contents
    let successCount = 0
    let failCount = 0
    const failedFiles: Array<{ name: string; reason: string }> = []

    for (const file of validFiles) {
        // Get file relative path (webkitRelativePath) to preserve subdirectory structure
        const relativePath = (file as any).webkitRelativePath
        let fileName = file.name
        if (relativePath) {
            const pathParts = relativePath.split('/')
            if (pathParts.length > 2) {
                const subPath = pathParts.slice(1, -1).join('/')
                fileName = `${subPath}/${file.name}`
            }
        }

        const uploadId = `${file.name}_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`
        let progress = 0
        let status: 'uploading' | 'success' | 'error' = 'uploading'
        let errorReason: string | undefined

        window.dispatchEvent(new CustomEvent('knowledgeFileUploadStart', {
            detail: {
                kbId,
                uploadId,
                fileName
            }
        }))

        try {
            // Get currently selected tag ID
            const tagIdToUpload = uiStore.selectedTagId !== '__untagged__' ? uiStore.selectedTagId : undefined
            await uploadKnowledgeFile(
                kbId,
                { file, fileName, tag_id: tagIdToUpload },
                (progressEvent: any) => {
                    if (progressEvent?.total) {
                        progress = Math.round((progressEvent.loaded * 100) / progressEvent.total)
                        window.dispatchEvent(new CustomEvent('knowledgeFileUploadProgress', {
                            detail: {
                                kbId,
                                uploadId,
                                progress
                            }
                        }))
                    }
                }
            )
            successCount++
            status = 'success'
            progress = 100
        } catch (error: any) {
            failCount++
            errorReason = error?.error?.message || error?.message || t('knowledgeBase.uploadFailed')
            if (error?.code === 'duplicate_file' || error?.error?.code === 'duplicate_file') {
                errorReason = t('knowledgeBase.fileExists')
            }
            failedFiles.push({ name: fileName, reason: errorReason })
            status = 'error'
        } finally {
            window.dispatchEvent(new CustomEvent('knowledgeFileUploadComplete', {
                detail: {
                    kbId,
                    uploadId,
                    status,
                    progress,
                    error: errorReason,
                    fileName
                }
            }))
        }
    }

    if (successCount > 0) {
        window.dispatchEvent(new CustomEvent('knowledgeFileUploaded', {
            detail: { kbId }
        }))
    }

    if (failCount === 0) {
        MessagePlugin.success(t('knowledgeBase.uploadAllSuccess', { count: successCount }))
    } else if (successCount > 0) {
        const { duplicateCount, nonDuplicateText } = summarizeFailedFiles(failedFiles)
        const extraSections: string[] = []
        if (duplicateCount > 0) {
            extraSections.push(t('knowledgeBase.duplicateFilesSkipped', { count: duplicateCount }))
        }
        if (nonDuplicateText) {
            extraSections.push(t('knowledgeBase.failedFilesList') + '\n' + nonDuplicateText)
        }
        const extraContent = extraSections.length ? '\n\n' + extraSections.join('\n\n') : ''
        MessagePlugin.warning({
            content: t('knowledgeBase.uploadPartialSuccess', {
                success: successCount,
                fail: failCount
            }) + extraContent,
            duration: 8000,
            closeBtn: true
        })
    } else {
        const { duplicateCount, nonDuplicateText } = summarizeFailedFiles(failedFiles)
        const extraSections: string[] = []
        if (duplicateCount > 0) {
            extraSections.push(t('knowledgeBase.duplicateFilesSkipped', { count: duplicateCount }))
        }
        if (nonDuplicateText) {
            extraSections.push(t('knowledgeBase.failedFilesList') + '\n' + nonDuplicateText)
        }
        const extraContent = extraSections.length ? '\n\n' + extraSections.join('\n\n') : ''
        MessagePlugin.error({
            content: t('knowledgeBase.uploadAllFailed') + extraContent,
            duration: 8000,
            closeBtn: true
        })
    }

    input.value = ''
}

const handleDocManualCreate = async () => {
    const kbId = await ensureDocKnowledgeBaseReady()
    if (!kbId) return
    uiStore.openManualEditor({
        mode: 'create',
        kbId,
        status: 'draft',
        onSuccess: ({ kbId: savedKbId }) => {
            if (savedKbId) {
                window.dispatchEvent(new CustomEvent('knowledgeFileUploaded', { detail: { kbId: savedKbId } }))
            }
        },
    })
}

const handleDocURLImport = async () => {
    const kbId = await ensureDocKnowledgeBaseReady()
    if (!kbId) return
    
    window.dispatchEvent(new CustomEvent('openURLImportDialog', {
        detail: { kbId }
    }))
}

const dispatchFaqMenuAction = (action: 'create' | 'import' | 'search' | 'export' | 'batch' | 'batchTag' | 'batchEnable' | 'batchDisable' | 'batchDelete', kbId: string) => {
    window.dispatchEvent(new CustomEvent('faqMenuAction', {
        detail: { action, kbId }
    }))
}

const handleFaqCreateFromMenu = async () => {
    const kbId = await getCurrentKbId()
    if (!kbId) {
        MessagePlugin.warning(t('knowledgeEditor.messages.missingId'))
        return
    }
    dispatchFaqMenuAction('create', kbId)
}

const handleFaqImportFromMenu = async () => {
    const kbId = await getCurrentKbId()
    if (!kbId) {
        MessagePlugin.warning(t('knowledgeEditor.messages.missingId'))
        return
    }
    dispatchFaqMenuAction('import', kbId)
}

const handleFaqSearchTestFromMenu = async () => {
    const kbId = await getCurrentKbId()
    if (!kbId) {
        MessagePlugin.warning(t('knowledgeEditor.messages.missingId'))
        return
    }
    dispatchFaqMenuAction('search', kbId)
}

const handleFaqExportFromMenu = async () => {
    const kbId = await getCurrentKbId()
    if (!kbId) {
        MessagePlugin.warning(t('knowledgeEditor.messages.missingId'))
        return
    }
    dispatchFaqMenuAction('export', kbId)
}

const faqBatchActionOptions = computed(() => {
  if (selectedFaqCount.value === 0) {
    return []
  }
  const options = [
    { 
      content: `${t('knowledgeEditor.faq.batchUpdateTag')} (${selectedFaqCount.value})`, 
      value: 'batchTag', 
      icon: 'folder'
    }
  ]
  
  // Show batch enable or disable based on selected items' state
  if (selectedFaqDisabledCount.value > 0) {
    options.push({
      content: `${t('knowledgeEditor.faq.batchEnable')} (${selectedFaqDisabledCount.value})`,
      value: 'batchEnable',
      icon: 'check-circle',
    })
  }
  if (selectedFaqEnabledCount.value > 0) {
    options.push({
      content: `${t('knowledgeEditor.faq.batchDisable')} (${selectedFaqEnabledCount.value})`,
      value: 'batchDisable',
      icon: 'close-circle',
    })
  }
  
  options.push({
    content: `${t('knowledgeEditor.faqImport.deleteSelected')} (${selectedFaqCount.value})`,
    value: 'batchDelete',
    icon: 'delete',
  })
  
  return options
})

const handleFaqBatchActionFromMenu = async (data: { value: string }) => {
  const kbId = await getCurrentKbId()
  if (!kbId) {
    MessagePlugin.warning(t('knowledgeEditor.messages.missingId'))
    return
  }
  if (selectedFaqCount.value === 0) {
    MessagePlugin.warning(t('knowledgeEditor.faq.selectEntriesFirst') || 'Please select FAQ entries to operate on first')
    return
  }
  dispatchFaqMenuAction(data.value as 'batchTag' | 'batchEnable' | 'batchDisable' | 'batchDelete', kbId)
}

const handleCreateKnowledgeBase = () => {
    uiStore.openCreateKB()
}

const handleCreateAgent = () => {
    // Trigger create agent event, handled by AgentList page listener
    window.dispatchEvent(new CustomEvent('openAgentEditor', {
        detail: { mode: 'create' }
    }))
}

</script>
<style lang="less" scoped>
.aside_box {
    min-width: 260px;
    padding: 8px;
    background: #fff;
    box-sizing: border-box;
    height: 100vh;
    overflow: hidden;
    display: flex;
    flex-direction: column;
    /* 与右侧内容区统一的细分界，减少割裂感 */
    border-right: 1px solid #e7ebf0;
    box-shadow: 1px 0 0 rgba(0, 0, 0, 0.02);

    .logo_box {
        height: 80px;
        display: flex;
        align-items: center;
        .logo{
            width: 134px;
            height: auto;
            margin-left: 24px;
        }
    }

    .logo_img {
        margin-left: 24px;
        width: 30px;
        height: 30px;
        margin-right: 7.25px;
    }

    .logo_txt {
        transform: rotate(0.049deg);
        color: #000000;
        font-family: "TencentSans";
        font-size: 24.12px;
        font-style: normal;
        font-weight: W7;
        line-height: 21.7px;
    }

    .menu_top {
        flex: 1;
        display: flex;
        flex-direction: column;
        overflow: hidden;
        min-height: 0;
    }

    .menu_bottom {
        flex-shrink: 0;
        display: flex;
        flex-direction: column;
    }

    .menu_box {
        display: flex;
        flex-direction: column;
        
        &.has-submenu {
            flex: 1;
            min-height: 0;
        }
    }


    .upload-file-wrap {
        padding: 6px;
        border-radius: 3px;
        height: 32px;
        width: 32px;
        box-sizing: border-box;
    }

    .upload-file-wrap:hover {
        background-color: #dbede4;
        color: #07C05F;

    }

    .upload-file-icon {
        width: 20px;
        height: 20px;
        color: rgba(0, 0, 0, 0.6);
    }

    .active-upload {
        color: #07C05F;
    }

    .menu_item_active {
        border-radius: 4px;
        background: #07c05f1a !important;

        .menu_icon,
        .menu_title {
            color: #07c05f !important;
        }

        .menu-create-hint {
            color: #07c05f !important;
            opacity: 1;
        }
    }

    .menu_item_c_active {

        .menu_icon,
        .menu_title {
            color: #000000e6;
        }
    }

    .menu_p {
        height: 56px;
        padding: 6px 0;
        box-sizing: border-box;
    }

    .menu_item {
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: space-between;
        height: 48px;
        padding: 13px 8px 13px 16px;
        box-sizing: border-box;
        margin-bottom: 4px;

        .menu_item-box {
            display: flex;
            align-items: center;
        }

        &:hover {
            border-radius: 4px;
            background: #30323605;
            color: #00000099;

            .menu_icon,
            .menu_title {
                color: #00000099;
            }
        }
    }

    .menu_icon {
        display: flex;
        margin-right: 10px;
        color: #00000099;

        .icon {
            width: 20px;
            height: 20px;
            fill: currentColor;
            overflow: hidden;
        }
    }

    .menu_title {
        color: #00000099;
        text-overflow: ellipsis;
        font-family: "PingFang SC";
        font-size: 14px;
        font-style: normal;
        font-weight: 600;
        line-height: 22px;
        overflow: hidden;
        white-space: nowrap;
        max-width: 120px;
        flex: 1;
    }

    .submenu {
        font-family: "PingFang SC";
        font-size: 14px;
        font-style: normal;
        overflow-y: auto;
        scrollbar-width: none;
        flex: 1;
        min-height: 0;
        margin-left: 4px;
    }
    
    .timeline_header {
        font-family: "PingFang SC";
        font-size: 12px;
        font-weight: 600;
        color: #00000066;
        padding: 12px 18px 6px 18px;
        margin-top: 8px;
        line-height: 20px;
        user-select: none;
        
        &:first-child {
            margin-top: 4px;
        }
    }

    .submenu_item_p {
        height: 44px;
        padding: 4px 0px 4px 0px;
        box-sizing: border-box;
    }


    .submenu_item {
        cursor: pointer;
        display: flex;
        align-items: center;
        color: #00000099;
        font-weight: 400;
        line-height: 22px;
        height: 36px;
        padding-left: 0px;
        padding-right: 14px;
        position: relative;

        .submenu_title {
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
        }

        .menu-more-wrap {
            margin-left: auto;
            opacity: 0;
            transition: opacity 0.2s ease;
        }

        .menu-more {
            display: inline-block;
            font-weight: bold;
            color: #07C05F;
        }

        .sub_title {
            margin-left: 14px;
        }

        &:hover {
            background: #30323605;
            color: #00000099;
            border-radius: 3px;

            .menu-more {
                color: #00000099;
            }

            .menu-more-wrap {
                opacity: 1;
            }

            .submenu_title {
                max-width: 160px !important;

            }
        }
    }

    .submenu_item_active {
        background: #07c05f1a !important;
        color: #07c05f !important;
        border-radius: 3px;

        .menu-more {
            color: #07c05f !important;
        }

        .menu-more-wrap {
            opacity: 1;
        }

        .submenu_title {
            max-width: 160px !important;
        }
    }
}

/* Knowledge base dropdown menu styles */
.kb-dropdown-icon {
    margin-left: auto;
    color: #666;
    transition: transform 0.3s ease, color 0.2s ease;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 16px;
    height: 16px;
    
    &.rotate-180 {
        transform: rotate(180deg);
    }
    
    &:hover {
        color: #07c05f;
    }
    
    &.active {
        color: #07c05f;
    }
    
    &.active:hover {
        color: #05a04f;
    }
    
    svg {
        width: 12px;
        height: 12px;
        transition: inherit;
    }
}

.kb-dropdown-menu {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    background: #fff;
    border: 1px solid #e5e7eb;
    border-radius: 6px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    z-index: 1000;
    max-height: 200px;
    overflow-y: auto;
}

.kb-dropdown-item {
    padding: 8px 16px;
    cursor: pointer;
    transition: background-color 0.2s ease;
    font-size: 14px;
    color: #333;
    
    &:hover {
        background-color: #f5f5f5;
    }
    
    &.active {
        background-color: #07c05f1a;
        color: #07c05f;
        font-weight: 500;
    }
    
    &:first-child {
        border-radius: 6px 6px 0 0;
    }
    
    &:last-child {
        border-radius: 0 0 6px 6px;
    }
}

.menu_item-box {
    display: flex;
    align-items: center;
    width: 100%;
    position: relative;
}

.menu-create-hint {
    margin-left: auto;
    margin-right: 8px;
    font-size: 16px;
    color: #07c05f;
    opacity: 0.7;
    transition: opacity 0.2s ease;
    flex-shrink: 0;
}

.menu_item:hover .menu-create-hint {
    opacity: 1;
}

.menu-pending-badge {
    min-width: 18px;
    height: 18px;
    padding: 0 5px;
    margin-left: 6px;
    border-radius: 9px;
    background: rgba(250, 173, 20, 0.2);
    color: #d48806;
    font-size: 12px;
    font-weight: 600;
    line-height: 18px;
    text-align: center;
    flex-shrink: 0;
}

.menu_box {
    position: relative;
}
</style>
<style lang="less">
// Upload operation dropdown menu styles - Global styles (because TDesign dropdown menus are mounted to body)
// Use more specific selectors to match upload operation dropdown menu
.t-popup[data-popper-placement^="right"] {
    .t-popup__content {
        .t-dropdown__menu {
            background: #ffffff !important;
            border: 1px solid #e5e7eb !important;
            border-radius: 6px !important;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1) !important;
            padding: 4px !important;
            min-width: 100px !important;
        }

        .t-dropdown__item {
            padding: 8px 12px !important;
            border-radius: 4px !important;
            margin: 2px 0 !important;
            transition: all 0.2s ease !important;
            font-size: 14px !important;
            color: #333333 !important;
            min-width: auto !important;
            max-width: none !important;
            width: auto !important;
            cursor: pointer !important;

            &:hover {
                background: #f5f7fa !important;
                color: #07c05f !important;
            }

            .t-dropdown__item-text {
                color: inherit !important;
                font-size: 14px !important;
                line-height: 20px !important;
                white-space: nowrap !important;
            }
        }
    }
}

// Logout confirmation box styles
:deep(.t-popconfirm) {
    .t-popconfirm__content {
        background: #fff;
        border: 1px solid #e7e7e7;
        border-radius: 6px;
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
        padding: 12px 16px;
        font-size: 14px;
        color: #333;
        max-width: 200px;
    }
    
    .t-popconfirm__arrow {
        border-bottom-color: #e7e7e7;
    }
    
    .t-popconfirm__arrow::after {
        border-bottom-color: #fff;
    }
    
    .t-popconfirm__buttons {
        margin-top: 8px;
        display: flex;
        justify-content: flex-end;
        gap: 8px;
    }
    
    .t-button--variant-outline {
        border-color: #d9d9d9;
        color: #666;
    }
    
    .t-button--theme-danger {
        background-color: #ff4d4f;
        border-color: #ff4d4f;
    }
    
    .t-button--theme-danger:hover {
        background-color: #ff7875;
        border-color: #ff7875;
    }
}
</style>