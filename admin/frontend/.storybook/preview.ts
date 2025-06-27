// .storybook/preview.ts
import type {Preview} from "@storybook/vue3";

const preview: Preview = {
    parameters: {
        controls: {
            matchers: {
                color: /(background|color)$/i,
                date: /Date$/,
            },
        },
        // Используем правильный формат для цветов
        backgrounds: {
            default: 'dark',
            values: [
                {name: 'dark', value: '#1a1a1a'}, // Используем HEX формат
                {name: 'light', value: '#ffffff'},
            ],
        },
        // Отключаем тему docs, чтобы использовать стандартную
        docs: {
            theme: undefined, // Или удалите этот параметр полностью
        },
    },
};

export default preview;
