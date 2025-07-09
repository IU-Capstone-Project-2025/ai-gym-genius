import type { Meta, StoryObj } from '@storybook/vue3'
import { UserManagementBoard } from '../index'

const meta = {
  title: 'Organisms/UserManagementBoard',
  component: UserManagementBoard,
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof UserManagementBoard>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {},
}

export const DarkMode: Story = {
  args: {},
  parameters: {
    backgrounds: { default: 'dark' },
  },
}