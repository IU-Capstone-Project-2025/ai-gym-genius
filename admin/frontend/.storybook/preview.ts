// .storybook/preview.ts
import type { Preview } from "@storybook/vue3";

const preview: Preview = {
  parameters: {
    // Удаляем блок actions, который вызывает предупреждение
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/,
      },
    },
  },
};

export default preview;
