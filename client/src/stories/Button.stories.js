import { Button } from '../components';

export default {
  title: 'Button',
  component: Button,
  argTypes: {
    color: {
      control: { type: 'select' },
      options: ['primary', 'gray', 'danger'],
    },
    size: {
      control: { type: 'select' },
      options: ['sm', 'md', 'lg'],
    },
    type: {
      control: { type: 'select' },
      options: ['button', 'reset', 'submit'],
    },
  },
};

const Template = (args) => ({
  components: { Button },

  setup() {
    return { args };
  },

  template: '<Button v-bind="args"/>',
});

const DefaultArgs = {
  value: 'button',
  name: 'button',
  text: 'button',
}

export const PrimarySmall = Template.bind({});
PrimarySmall.args = {
  ...DefaultArgs,
  color: 'primary',
  size: 'sm',
  type: 'button',
};

export const PrimaryMedium = Template.bind({});
PrimaryMedium.args = {
  ...DefaultArgs,
  color: 'primary',
  size: 'md',
  type: 'button',
};

export const PrimaryLarge = Template.bind({});
PrimaryLarge.args = {
  ...DefaultArgs,
  color: 'primary',
  size: 'lg',
  type: 'button',
};

export const GraySmall = Template.bind({});
GraySmall.args = {
  ...DefaultArgs,
  color: 'gray',
  size: 'sm',
  type: 'button',
};

export const GrayMedium = Template.bind({});
GrayMedium.args = {
  ...DefaultArgs,
  color: 'gray',
  size: 'md',
  type: 'button',
};

export const GrayLarge = Template.bind({});
GrayLarge.args = {
  ...DefaultArgs,
  color: 'gray',
  size: 'lg',
  type: 'button',
};

export const DangerSmall = Template.bind({});
DangerSmall.args = {
  ...DefaultArgs,
  color: 'danger',
  size: 'sm',
  type: 'button',
};

export const DangerMedium = Template.bind({});
DangerMedium.args = {
  ...DefaultArgs,
  color: 'danger',
  size: 'md',
  type: 'button',
};

export const DangerLarge = Template.bind({});
DangerLarge.args = {
  ...DefaultArgs,
  color: 'danger',
  size: 'lg',
  type: 'button',
};
