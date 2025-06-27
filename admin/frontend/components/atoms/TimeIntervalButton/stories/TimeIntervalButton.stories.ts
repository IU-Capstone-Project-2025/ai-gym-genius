// components/atoms/TimeIntervalButton/stories/TimeIntervalButton.stories.ts
import type { Meta, StoryObj } from '@storybook/vue3';
import TimeIntervalButton from "../ui/TimeIntervalButton.vue";

// Без прямого импорта UButton
const meta = {
    title: "Atoms/TimeIntervalButton",
    component: TimeIntervalButton,
    tags: ['autodocs'],
    argTypes: {
        label: { control: 'text' },
        isActive: { control: 'boolean' },
        select: { action: 'selected' }
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