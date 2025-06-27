// components/atoms/TimeIntervalButton/stories/TimeIntervalButton.stories.ts
import type {Meta, StoryObj} from '@storybook/vue3';
import TimeIntervalButton from "../ui/TimeIntervalButton.vue";

// Без прямого импорта UButton
const meta = {
    title: "Atoms/TimeIntervalButton",
    component: TimeIntervalButton,
    tags: ['autodocs'],
    argTypes: {
        label: {control: 'text'},
        isActive: {control: 'boolean'},
        select: {action: 'selected'}
    }
} satisfies Meta<typeof TimeIntervalButton>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Active: Story = {
    args: {
        label: '30 min',
        isActive: true
    }
};

export const Inactive: Story = {
    args: {
        label: '1 hour',
        isActive: false
    }
};

// Добавьте эту историю в существующий файл
export const DarkMode: Story = {
    render: () => ({
        components: {TimeIntervalButton},
        setup() {
            return {
                args: {
                    label: '30 min',
                    isActive: true
                }
            };
        },
        template: `
          <div style="padding: 2rem; background-color: #1a1a1a;" class="dark">
            <div class="flex space-x-4">
              <TimeIntervalButton label="30 min" :is-active="true" @select="() => {}"/>
              <TimeIntervalButton label="1 hour" :is-active="false" @select="() => {}"/>
            </div>
          </div>
        `
    })
};
