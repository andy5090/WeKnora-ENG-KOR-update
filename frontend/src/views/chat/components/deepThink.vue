<template>
    <div class='deep-think'>
        <div class="think-header" @click="toggleFold">
            <div class="think-title">
                <span v-if="deepSession.thinking" class="thinking-status">
                    <img class="thinking-gif" src="@/assets/img/think.gif" :alt="$t('chat.thinkingAlt')">
                    <span class="thinking-text">{{ $t('chat.thinking') }}</span>
                </span>
                <span v-else class="done-status">
                    <img class="done-icon" src="@/assets/img/Frame3718.svg" :alt="$t('chat.deepThoughtAlt')">
                    <span class="done-text">{{ $t('chat.deepThoughtCompleted') }}</span>
                </span>
            </div>
            <div class="toggle-icon-wrapper">
                <t-icon :name="isFold ? 'chevron-down' : 'chevron-up'" class="toggle-icon" />
            </div>
        </div>
        <div class="think-content" v-show="!isFold || deepSession.thinking">
            <div ref="contentInnerRef" class="content-inner" v-html="safeProcessThinkContent(deepSession.thinkContent)"></div>
        </div>
    </div>
</template>
<script setup>
import { watch, ref, defineProps, onMounted, nextTick } from 'vue';
import { sanitizeHTML } from '@/utils/security';
import { useI18n } from 'vue-i18n';

const isFold = ref(false)
const contentInnerRef = ref(null)
const { t } = useI18n()
const props = defineProps({
    // Required field
    deepSession: {
        type: Object,
        required: false
    }
});

// Check on initialization: if thinking is completed (loaded from history), collapse by default
onMounted(() => {
    if (props.deepSession?.thinking === false) {
        isFold.value = true;
    }
});

// Watch thinking state changes, auto-collapse
watch(
    () => props.deepSession?.thinking,
    (newVal, oldVal) => {
        // When thinking changes from true to false, auto-collapse thinking content
        // Only trigger in streaming output scenario (oldVal is true)
        if (oldVal === true && newVal === false) {
            isFold.value = true;
        }
    }
);

// Watch content changes, auto-scroll to bottom
watch(
    () => props.deepSession?.thinkContent,
    () => {
        // Only scroll when thinking is in progress
        if (props.deepSession?.thinking) {
            nextTick(() => {
                if (contentInnerRef.value) {
                    contentInnerRef.value.scrollTop = contentInnerRef.value.scrollHeight;
                }
            });
        }
    }
);

const toggleFold = () => {
    // Only can collapse/expand after thinking is completed
    if (!props.deepSession?.thinking) {
        isFold.value = !isFold.value;
    }
}

// Safely process thinking content to prevent XSS attacks
const safeProcessThinkContent = (content) => {
    if (!content || typeof content !== 'string') return '';
    
    // First handle line breaks
    const contentWithBreaks = content.replace(/\n/g, '<br/>');
    
    // Use DOMPurify for security cleanup, allow basic text formatting tags
    const cleanContent = sanitizeHTML(contentWithBreaks);
    
    return cleanContent;
};
</script>
<style lang="less" scoped>
.deep-think {
    display: flex;
    flex-direction: column;
    font-size: 12px;
    width: 100%;
    border-radius: 8px;
    background-color: #ffffff;
    box-shadow: 0 2px 4px rgba(7, 192, 95, 0.08);
    overflow: hidden;
    box-sizing: border-box;
    transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
    margin: -8px 0px 10px 0px;

    .think-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 6px 14px;
        color: #333333;
        font-weight: 500;
        cursor: pointer;
        user-select: none;

        &:hover {
            background-color: rgba(7, 192, 95, 0.04);
        }

        .think-title {
            display: flex;
            align-items: center;
        }

        .thinking-status {
            display: flex;
            align-items: center;
            
            .thinking-gif {
                width: 16px;
                height: 16px;
                margin-right: 8px;
            }
            
            .thinking-text {
                font-size: 12px;
                color: #333333;
                white-space: nowrap;
            }
        }

        .done-status {
            display: flex;
            align-items: center;
            
            .done-icon {
                width: 16px;
                height: 16px;
                margin-right: 8px;
            }
            
            .done-text {
                font-size: 12px;
                color: #333333;
                white-space: nowrap;
            }
        }

        .toggle-icon-wrapper {
            font-size: 14px;
            padding: 0 2px 1px 2px;
            color: #07c05f;
            
            .toggle-icon {
                transition: transform 0.2s;
            }
        }
    }

    .think-content {
        border-top: 1px solid #f0f0f0;
        
        .content-inner {
            padding: 8px 14px;
            font-size: 12px;
            line-height: 1.6;
            color: #666666;
            max-height: 200px;
            overflow-y: auto;
            word-break: break-word;
            
            &::-webkit-scrollbar {
                width: 4px;
            }
            
            &::-webkit-scrollbar-thumb {
                background: rgba(0, 0, 0, 0.1);
                border-radius: 2px;
            }
        }
    }
}
</style>
