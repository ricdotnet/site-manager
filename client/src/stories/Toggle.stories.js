import { Toggle } from '../components';

export default {
  title: 'Toggle',
  component: Toggle,
};

const Template = (args) => ({
  components: { Toggle },

  setup() {
    return { args };
  },

  template: '<Toggle v-bind="args"/>',
});

export const ToggleDefault = Template.bind({});
ToggleDefault.args = {
  name: 'toggle',
  isChecked: false,
};

export const ToggleTheme = Template.bind({});
ToggleTheme.args = {
  name: 'toggle-theme',
  isChecked: false,
  isThemeToggle: true,
};
