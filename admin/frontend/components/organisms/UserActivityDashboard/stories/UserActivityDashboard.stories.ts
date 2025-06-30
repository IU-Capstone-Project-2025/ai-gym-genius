// components/organisms/UserActivityCard/stories/UserActivityCard.stories.ts
import type {Meta, StoryObj} from '@storybook/vue3';
import UserActivityCard from "../ui/UserActivityDashboard.vue";
import {ref, computed} from 'vue';

// Mock the useUserActivity composable
const mockUserActivity = (isLoadingState = false, intervalState = '24h', dataState = {
    labels: ['9AM', '10AM', '11AM', '12PM', '1PM', '2PM', '3PM', '4PM', '5PM'],
    values: [45, 52, 49, 60, 72, 56, 42, 48, 53]
}) => {
    // Return the mocked version in window object to be available in component
    window.useUserActivity = () => {
        const isLoading = ref(isLoadingState);
        const selectedInterval = ref(intervalState);
        const activityData = ref(dataState);

        const fetchActivityData = () => {
            isLoading.value = true;
            // Simulate API call
            setTimeout(() => {
                isLoading.value = false;
            }, 1500);
        };

        return {
            isLoading,
            selectedInterval,
            activityData,
            fetchActivityData
        };
    };
};

const meta = {
    title: "Organisms/UserActivityCard",
    component: UserActivityCard,
    tags: ['autodocs'],
    parameters: {
        componentSubtitle: 'Card showing user activity data with time interval selection',
    }
} satisfies Meta<typeof UserActivityCard>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
    render: () => ({
        components: {UserActivityCard},
        setup() {
            mockUserActivity(false);
            return {};
        },
        template: '<div style="max-width: 800px; margin: 0 auto;"><UserActivityCard /></div>'
    })
};

export const Loading: Story = {
    render: () => ({
        components: {UserActivityCard},
        setup() {
            mockUserActivity(true);
            return {};
        },
        template: '<div style="max-width: 800px; margin: 0 auto;"><UserActivityCard /></div>'
    })
};

export const WeeklyData: Story = {
    render: () => ({
        components: {UserActivityCard},
        setup() {
            mockUserActivity(false, '7d', {
                labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
                values: [120, 150, 180, 145, 190, 75, 90]
            });
            return {};
        },
        template: '<div style="max-width: 800px; margin: 0 auto;"><UserActivityCard /></div>'
    })
};

export const MonthlyData: Story = {
    render: () => ({
        components: {UserActivityCard},
        setup() {
            mockUserActivity(false, '30d', {
                labels: ['Week 1', 'Week 2', 'Week 3', 'Week 4'],
                values: [450, 520, 480, 550]
            });
            return {};
        },
        template: '<div style="max-width: 800px; margin: 0 auto;"><UserActivityCard /></div>'
    })
};

export const LowActivity: Story = {
    render: () => ({
        components: {UserActivityCard},
        setup() {
            mockUserActivity(false, '24h', {
                labels: ['9AM', '10AM', '11AM', '12PM', '1PM', '2PM', '3PM', '4PM', '5PM'],
                values: [5, 8, 3, 7, 10, 6, 4, 9, 7]
            });
            return {};
        },
        template: '<div style="max-width: 800px; margin: 0 auto;"><UserActivityCard /></div>'
    })
};

export const HighActivity: Story = {
    render: () => ({
        components: {UserActivityCard},
        setup() {
            mockUserActivity(false, '24h', {
                labels: ['9AM', '10AM', '11AM', '12PM', '1PM', '2PM', '3PM', '4PM', '5PM'],
                values: [245, 352, 289, 360, 472, 356, 342, 448, 353]
            });
            return {};
        },
        template: '<div style="max-width: 800px; margin: 0 auto;"><UserActivityCard /></div>'
    })
};

export const DarkMode: Story = {
    render: () => ({
        components: {UserActivityCard},
        setup() {
            mockUserActivity(false);
            return {};
        },
        template: '<div style="max-width: 800px; margin: 0 auto; background-color: #1a1a1a; padding: 20px;" class="dark"><UserActivityCard /></div>'
    })
};
