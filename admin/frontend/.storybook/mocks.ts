// .storybook/mocks.ts
import { ComponentStory } from '@storybook/vue3';

export const mockNuxtLink = (story: ComponentStory): any => ({
    components: { story },
    template: `
    <story />
  `,
    setup() {
        // This will be called before each story render
        (window as any).NuxtLink = 'a'; // Replace NuxtLink with 'a' tag in stories
        return {};
    },
});
