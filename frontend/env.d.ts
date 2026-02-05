/// <reference types="vite/client" />
// Configure this file to resolve error: Cannot find module "@/views/login/index.vue" or its corresponding type declarations. ts(2307)
// This code tells TypeScript that all files ending with .vue are Vue components and can be imported via import statements. This usually resolves module recognition issues.
declare module '*.vue' {
    import { Component } from 'vue'; const component: Component; export default component;
}
