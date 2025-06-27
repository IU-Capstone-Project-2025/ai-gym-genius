// components/organisms/AuthForm/stories/AuthForm.stories.ts
import type {Meta, StoryObj} from '@storybook/vue3';
import AuthForm from "../ui/AuthForm.vue";

const meta = {
    title: "Organisms/AuthForm",
    component: AuthForm,
    tags: ['autodocs'],
    argTypes: {
        type: {
            control: 'select',
            options: ['sign_in', 'sign_up'],
            description: 'Authentication form mode'
        },
        switch: {
            action: 'switch',
            description: 'Event emitted when switching between sign-in and sign-up modes'
        },
        action: {
            action: 'action',
            description: 'Event emitted when submitting the form with login and password'
        }
    },
    parameters: {
        componentSubtitle: 'Authentication form for user login and registration',
        layout: 'fullscreen'
    }
} satisfies Meta<typeof AuthForm>;

export default meta;
type Story = StoryObj<typeof meta>;

export const LoginMode: Story = {
    args: {
        type: 'sign_in'
    }
};

export const SignupMode: Story = {
    args: {
        type: 'sign_up'
    }
};

// Interactive version that shows form behavior
export const Interactive: Story = {
    parameters: {
        docs: {
            description: {
                story: 'Interactive example showing how the form switches between login and signup modes'
            }
        }
    },
    render: (args) => ({
        components: {AuthForm},
        setup() {
            const formType = ref(args.type || 'sign_in');

            function handleSwitch(type) {
                formType.value = type;
                args.switch(type);
            }

            function handleAction(login, password) {
                args.action(login, password);
            }

            return {formType, handleSwitch, handleAction};
        },
        template: `
          <div class="h-screen bg-gray-100 p-4">
            <AuthForm
                :type="formType"
                @switch="handleSwitch"
                @action="handleAction"
            />
            <div class="fixed bottom-4 left-4 p-4 bg-white rounded shadow">
              <p class="text-sm font-semibold mb-2">Debug Info:</p>
              <p class="text-xs">Current Mode: {{ formType }}</p>
              <p class="text-xs mt-2">Try clicking the Log in/Sign up buttons to switch modes</p>
              <p class="text-xs mt-2">Fill out both fields and click the active button to trigger the action event</p>
            </div>
          </div>
        `
    })
};

// Story that demonstrates form validation behavior
export const FormValidation: Story = {
    args: {
        type: 'sign_in'
    },
    parameters: {
        docs: {
            description: {
                story: 'Demonstrates that the action event is only triggered when both fields are filled'
            }
        }
    },
    render: (args) => ({
        components: {AuthForm},
        setup() {
            const actionCalled = ref(false);
            const actionDetails = ref(null);

            function handleAction(login, password) {
                actionCalled.value = true;
                actionDetails.value = {login, password};
                args.action(login, password);

                // Reset after 2 seconds
                setTimeout(() => {
                    actionCalled.value = false;
                    actionDetails.value = null;
                }, 2000);
            }

            return {args, actionCalled, actionDetails, handleAction};
        },
        template: `
          <div class="h-screen bg-gray-100 p-4">
            <AuthForm
                :type="args.type"
                @switch="args.switch"
                @action="handleAction"
            />
            <div class="fixed bottom-4 left-4 p-4 bg-white rounded shadow">
              <p class="text-sm font-semibold mb-2">Form Submission:</p>
              <div v-if="actionCalled" class="text-xs p-2 bg-green-100 rounded">
                <p class="font-semibold text-green-700">Form Submitted Successfully!</p>
                <p class="mt-1">Login: {{ actionDetails.login }}</p>
                <p>Password: {{ actionDetails.password }}</p>
              </div>
              <p v-else class="text-xs">Fill both fields and click the main button to submit</p>
            </div>
          </div>
        `
    })
};

// Добавьте эту историю в существующий файл
export const DarkMode: Story = {
    render: () => ({
        components: { AuthForm },
        setup() {
            const formType = ref('sign_in');

            function handleSwitch(type) {
                formType.value = type;
            }

            function handleAction(login, password) {
                console.log('Action triggered:', login, password);
            }

            return { formType, handleSwitch, handleAction };
        },
        template: `
            <div style="min-height: 100vh; background-color: #1a1a1a; padding: 2rem;" class="dark">
                <AuthForm 
                    :type="formType"
                    @switch="handleSwitch"
                    @action="handleAction"
                />
            </div>
        `
    })
};
