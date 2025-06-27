// components/molecules/TimeIntervalSelector/stories/TimeIntervalSelector.stories.ts
import type {Meta, StoryObj} from '@storybook/vue3';
import TimeIntervalSelector from "../ui/TimeIntervalSelector.vue";

const meta = {
    title: "Molecules/TimeIntervalSelector",
    component: TimeIntervalSelector,
    tags: ['autodocs'],
    argTypes: {
        selectedInterval: {
            control: 'select',
            options: ['24h', '7d', '30d', '90d'],
            description: 'Currently selected time interval'
        },
        'update:selectedInterval': {
            action: 'update:selectedInterval',
            description: 'Event emitted when interval selection changes'
        }
    },
    parameters: {
        componentSubtitle: 'A selector for time interval filtering',
    }
} satisfies Meta<typeof TimeIntervalSelector>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
    args: {
        selectedInterval: '24h'
    }
};

export const SevenDaysSelected: Story = {
    args: {
        selectedInterval: '7d'
    }
};

export const ThirtyDaysSelected: Story = {
    args: {
        selectedInterval: '30d'
    }
};

export const NinetyDaysSelected: Story = {
    args: {
        selectedInterval: '90d'
    }
};

// Example with a decorator to show reactivity
export const Interactive: Story = {
    args: {
        selectedInterval: '24h'
    },
    parameters: {
        docs: {
            description: {
                story: 'Interactive example that allows changing the selected interval'
            }
        }
    },
    render: (args) => ({
        components: {TimeIntervalSelector},
        setup() {
            return {args};
        },
        template: `
          <div>
            <p class="mb-2">Selected: {{ args.selectedInterval }}</p>
            <TimeIntervalSelector
                :selectedInterval="args.selectedInterval"
                @update:selectedInterval="args.selectedInterval = $event; args['update:selectedInterval']($event)"
            />
          </div>
        `
    })
};

// Добавьте эту историю в существующий файл
export const DarkMode: Story = {
    render: () => ({
        components: {TimeIntervalSelector},
        setup() {
            const selectedInterval = ref('24h');
            return {selectedInterval};
        },
        template: `
          <div style="padding: 2rem; background-color: #1a1a1a;" class="dark">
            <TimeIntervalSelector
                :selectedInterval="selectedInterval"
                @update:selectedInterval="selectedInterval = $event"
            />
            <div class="mt-4 text-white">Selected: {{ selectedInterval }}</div>
          </div>
        `
    })
};

