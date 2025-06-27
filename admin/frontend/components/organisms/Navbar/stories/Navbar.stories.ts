// components/organisms/Navbar/stories/Navbar.stories.ts
import type { Meta, StoryObj } from '@storybook/vue3';
import Navbar from "../ui/Navbar.vue";

const meta = {
    title: "Organisms/Navbar",
    component: Navbar,
    tags: ['autodocs'],
    parameters: {
        layout: 'fullscreen',
        // Удалите любые параметры backgrounds здесь
    }
} satisfies Meta<typeof Navbar>;

export default meta;
type Story = StoryObj<typeof meta>;

// Базовая история с видимой кнопкой логина
export const WithLoginButton: Story = {
    render: () => ({
        components: { Navbar },
        setup() {
            window.useRoute = () => ({ name: 'dashboard' });
            return {};
        },
        // Используем класс вместо фона
        template: '<div style="min-height: 100vh; background-color: #1a1a1a;"><Navbar /></div>'
    })
};

// История без кнопки логина (на странице auth)
export const WithoutLoginButton: Story = {
    render: () => ({
        components: { Navbar },
        setup() {
            window.useRoute = () => ({ name: 'auth' });
            return {};
        },
        template: '<div style="min-height: 100vh; background-color: #1a1a1a;"><Navbar /></div>'
    })
};
