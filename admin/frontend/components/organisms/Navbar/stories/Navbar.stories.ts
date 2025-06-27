// components/organisms/Navbar/stories/Navbar.stories.ts
import type {Meta, StoryObj} from '@storybook/vue3';
import Navbar from "../ui/Navbar.vue";

const meta = {
    title: "Organisms/Navbar",
    component: Navbar,
    tags: ['autodocs'],
    parameters: {
        layout: 'fullscreen',
    }
} satisfies Meta<typeof Navbar>;

export default meta;
type Story = StoryObj<typeof meta>;

// Базовая история с видимой кнопкой логина
export const WithLoginButton: Story = {
    render: () => ({
        components: {Navbar},
        setup() {
            // Имитация нахождения не на странице auth
            window.useRoute = () => ({name: 'dashboard'});
            return {};
        },
        template: '<Navbar />'
    })
};

// История без кнопки логина (на странице auth)
export const WithoutLoginButton: Story = {
    render: () => ({
        components: {Navbar},
        setup() {
            // Имитация нахождения на странице auth
            window.useRoute = () => ({name: 'auth'});
            return {};
        },
        template: '<Navbar />'
    })
};

// Темная тема с кнопкой логина
export const DarkMode: Story = {
    render: () => ({
        components: {Navbar},
        setup() {
            // Глобально подменяем useRoute только для этой истории
            window.useRoute = () => ({
                name: 'dashboard' // Чтобы показать кнопку логина
            });
            return {};
        },
        template: `
          <div style="min-height: 100vh; background-color: #1a1a1a;" class="dark">
            <Navbar/>
          </div>
        `
    }),
    parameters: {
        backgrounds: {default: 'dark'}
    }
};

// Мобильная версия
export const Mobile: Story = {
    render: () => ({
        components: {Navbar},
        setup() {
            window.useRoute = () => ({name: 'dashboard'});
            return {};
        },
        template: '<Navbar />'
    }),
    parameters: {
        viewport: {
            defaultViewport: 'mobile1'
        }
    }
};

// Планшетная версия
export const Tablet: Story = {
    render: () => ({
        components: {Navbar},
        setup() {
            window.useRoute = () => ({name: 'dashboard'});
            return {};
        },
        template: '<Navbar />'
    }),
    parameters: {
        viewport: {
            defaultViewport: 'tablet'
        }
    }
};
