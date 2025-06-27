// components/atoms/ActivityChart/stories/ActivityChart.stories.ts
import type {Meta, StoryObj} from '@storybook/vue3';
import ActivityChart from "../ui/ActivityChart.vue";

const meta = {
    title: "Atoms/ActivityChart",
    component: ActivityChart,
    tags: ['autodocs'],
    argTypes: {
        data: {
            control: 'object',
            description: 'Chart data containing labels and values arrays'
        },
        title: {
            control: 'text',
            description: 'Optional chart title'
        }
    },
    parameters: {
        componentSubtitle: 'A line chart component to display user activity data',
    }
} satisfies Meta<typeof ActivityChart>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
    args: {
        data: {
            labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
            values: [65, 59, 80, 81, 56, 55, 40]
        },
        title: 'Weekly User Activity'
    }
};

export const WithoutTitle: Story = {
    args: {
        data: {
            labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
            values: [65, 59, 80, 81, 56, 55, 40]
        }
    }
};

export const HourlyData: Story = {
    args: {
        data: {
            labels: ['9AM', '10AM', '11AM', '12PM', '1PM', '2PM', '3PM', '4PM', '5PM'],
            values: [20, 35, 45, 30, 25, 60, 75, 65, 45]
        },
        title: 'Today\'s Hourly Activity'
    }
};

export const MonthlyData: Story = {
    args: {
        data: {
            labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
            values: [110, 130, 150, 175, 200, 225, 250, 260, 240, 200, 180, 160]
        },
        title: 'Monthly Active Users'
    }
};

export const FluctuatingData: Story = {
    args: {
        data: {
            labels: ['Week 1', 'Week 2', 'Week 3', 'Week 4', 'Week 5', 'Week 6'],
            values: [120, 60, 230, 90, 180, 75]
        },
        title: 'Fluctuating User Engagement'
    }
};
