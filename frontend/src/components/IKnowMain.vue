<template>
  <a-layout class="layout">
    <a-layout-header>
<!--      <div class="logo" />-->
      <a-menu
          theme="dark"
          mode="horizontal"
          v-model:selectedKeys="selectedKeys"
          :style="{ lineHeight: '64px' }"
      >
        <a-menu-item key="1" v-on:click="SwitchKeys">表格</a-menu-item>
        <a-menu-item key="2" v-on:click="SwitchKeys">思维导图</a-menu-item>
        <a-menu-item key="3" v-on:click="SwitchKeys">关于我</a-menu-item>
      </a-menu>
    </a-layout-header>
    <a-layout-content style="padding: 0 0px">
<!--      <a-breadcrumb style="margin: 16px 0">-->
<!--        &lt;!&ndash;<a-breadcrumb-item>Home</a-breadcrumb-item>&ndash;&gt;-->
<!--        &lt;!&ndash;<a-breadcrumb-item>List</a-breadcrumb-item>&ndash;&gt;-->
<!--        &lt;!&ndash;<a-breadcrumb-item>App</a-breadcrumb-item>&ndash;&gt;-->
<!--      </a-breadcrumb>-->
<!--      <div :style="{ background: '#fff', padding: '24px', minHeight: '280px' }">Content</div>-->
      <NodeTables></NodeTables>
    </a-layout-content>
    <a-layout-footer style="text-align: center">
      Ant Design ©2018 Created by Ant UED
    </a-layout-footer>
  </a-layout>
</template>
<script>
import { defineComponent, ref } from 'vue';
import NodeTables from "@/components/NodeTables";
export default defineComponent({
  components: {NodeTables},
  setup() {
    return {
      selectedKeys: ref(['1']),
    };
  },
  methods: {
  SwitchKeys: function () {
  }
  }
});
</script>
<style>
.site-layout-content {
  min-height: 280px;
  padding: 24px;
  background: #fff;
}
#components-layout-demo-top .logo {
  float: left;
  width: 120px;
  height: 31px;
  margin: 16px 24px 16px 0;
  background: rgba(255, 255, 255, 0.3);
}
.ant-row-rtl #components-layout-demo-top .logo {
  float: right;
  margin: 16px 0 16px 24px;
}

[data-theme='dark'] .site-layout-content {
  background: #141414;
}
</style>
