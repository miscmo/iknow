<template>
  <a-layout style="min-height: 100vh">
<!--      <div class="logo" />-->
    <a-layout-sider width="200" style="background: #fff">
<!--      <a-menu theme="dark" v-model:selectedKeys="selectedKeys" mode="inline">-->
<!--        <a-menu-item key="1">-->
<!--          <pie-chart-outlined />-->
<!--          <span>Option 1</span>-->
<!--        </a-menu-item>-->
<!--        <a-menu-item key="2">-->
<!--          <desktop-outlined />-->
<!--          <span>Option 2</span>-->
<!--        </a-menu-item>-->
<!--        <a-sub-menu key="sub1">-->
<!--          <template #title>-->
<!--            <span>-->
<!--              <user-outlined />-->
<!--              <span>User</span>-->
<!--            </span>-->
<!--          </template>-->
<!--          <a-menu-item key="3">Tom</a-menu-item>-->
<!--          <a-menu-item key="4">Bill</a-menu-item>-->
<!--          <a-menu-item key="5">Alex</a-menu-item>-->
<!--        </a-sub-menu>-->
<!--        <a-sub-menu key="sub2">-->
<!--          <template #title>-->
<!--            <span>-->
<!--              <team-outlined />-->
<!--              <span>Team</span>-->
<!--            </span>-->
<!--          </template>-->
<!--          <a-menu-item key="6">Team 1</a-menu-item>-->
<!--          <a-menu-item key="8">Team 2</a-menu-item>-->
<!--        </a-sub-menu>-->
<!--        <a-menu-item key="9">-->
<!--          <file-outlined />-->
<!--          <span>File</span>-->
<!--        </a-menu-item>-->
<!--      </a-menu>-->
      <div style="width: 256px">
        <a-button type="primary" @click="toggleCollapsed" style="margin-bottom: 16px">
          <MenuUnfoldOutlined v-if="collapsed" />
          <MenuFoldOutlined v-else />
        </a-button>
        <a-menu
            mode="inline"
            theme="dark"
            :inline-collapsed="collapsed"
            v-model:openKeys="openKeys"
            v-model:selectedKeys="selectedKeys"
        >
          <a-menu-item key="1">
            <template #icon>
              <PieChartOutlined />
            </template>
            <span>Option 1</span>
          </a-menu-item>
          <a-menu-item key="2">
            <template #icon>
              <DesktopOutlined />
            </template>
            <span>Option 2</span>
          </a-menu-item>
          <a-menu-item key="3">
            <template #icon>
              <InboxOutlined />
            </template>
            <span>Option 3</span>
          </a-menu-item>
          <a-sub-menu key="sub1">
            <template #icon>
              <MailOutlined />
            </template>
            <template #title>Navigation One</template>
            <a-menu-item key="5">Option 5</a-menu-item>
            <a-menu-item key="6">Option 6</a-menu-item>
            <a-menu-item key="7">Option 7</a-menu-item>
            <a-menu-item key="8">Option 8</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="sub2">
            <template #icon>
              <AppstoreOutlined />
            </template>
            <template #title>Navigation Two</template>
            <a-menu-item key="9">Option 9</a-menu-item>
            <a-menu-item key="10">Option 10</a-menu-item>
            <a-sub-menu key="sub3" title="Submenu">
              <a-menu-item key="11">Option 11</a-menu-item>
              <a-menu-item key="12">Option 12</a-menu-item>
            </a-sub-menu>
          </a-sub-menu>
        </a-menu>
      </div>
    </a-layout-sider>
    <a-layout>
<!--      <a-layout-header style="background: #fff; padding: 0" />-->
      <a-layout-content style="margin: 0 16px">
        <div align="right">
          <a-button v-on:click="SearchNodes">查询</a-button>
          <a-button v-on:click="AddNode">添加</a-button>

        </div>
<!--        <a-breadcrumb style="margin: 16px 0">-->
<!--          <a-breadcrumb-item>User</a-breadcrumb-item>-->
<!--          <a-breadcrumb-item>Bill</a-breadcrumb-item>-->
<!--        </a-breadcrumb>-->
<!--        <div :style="{ padding: '24px', background: '#fff', minHeight: '360px' }">-->
<!--          Bill is a cat.-->
<!--        </div>-->
        <a-table :columns="columns" :data-source="data" :row-selection="rowSelection" />
      </a-layout-content>
    </a-layout>
  </a-layout>


  <div>
  </div>
</template>


<script>

// import { defineComponent } from 'vue';
// import { defineComponent, ref } from 'vue';
import { defineComponent, reactive, toRefs, watch, ref } from 'vue';
import {
  PieChartOutlined,
  DesktopOutlined,
  // UserOutlined,
  // TeamOutlined,
  // FileOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  MailOutlined,
  InboxOutlined,
  AppstoreOutlined,
} from '@ant-design/icons-vue';
import axios from "axios";
const columns = [ { title: 'Id', dataIndex: 'id', key: 'id', },
    { title: '名称', dataIndex: 'name', key: 'name', /* width: '12%',*/ },
    { title: '描述', dataIndex: 'desc', key: 'desc', },
    { title: "创建时间", dataIndex: 'createTime', key: 'createTime', },
    { title: "更新时间", dataIndex: 'updateTime', key: 'updateTime', },
    { title: "类型", dataIndex: 'type', key: 'type', },
    { title: "子结点", dataIndex: 'subNodes', key: 'subNodes', },
    { title: "分数", dataIndex: 'score', key: 'score', },
    { title: "笔记", dataIndex: 'notes', key: 'notes', },
    { title: "分数", dataIndex: 'score', key: 'score', },
];
var data = [ ];

// let page = 1;
// const pageSize = 20;
const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
  },
  onSelect: (record, selected, selectedRows) => {
    console.log(record, selected, selectedRows);
  },
  onSelectAll: (selected, selectedRows, changeRows) => {
    console.log(selected, selectedRows, changeRows);
  },
};
export default defineComponent({
  name: "NodeTables",
  // setup() {
  //   return {
  //     //this.SearchNodes("", page, pageSize),
  //     data,
  //     columns,
  //     rowSelection,
  //   };
  // },
  components: {
    PieChartOutlined,
    DesktopOutlined,
    // UserOutlined,
    // TeamOutlined,
    // FileOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    MailOutlined,
    InboxOutlined,
    AppstoreOutlined
  },
  data() {
    this.SearchNodes()
    return {
      data,
      columns,
      rowSelection,
      // collapsed: ref(false),
      selectedKeys: ref(['1']),
    }
  },
      setup() {
        const state = reactive({
          collapsed: false,
          selectedKeys: ['1'],
          openKeys: ['sub1'],
          preOpenKeys: ['sub1'],
        });
        watch(
            () => state.openKeys,
            (val, oldVal) => {
              state.preOpenKeys = oldVal;
            },
        );

        const toggleCollapsed = () => {
          state.collapsed = !state.collapsed;
          state.openKeys = state.collapsed ? [] : state.preOpenKeys;
        };

        return { ...toRefs(state), toggleCollapsed };
      },


  methods:{
    SearchNodes: function () {
      var searchReq = {"id": 0, "page": 1, "pageSize": 10}
      /*eslint-disable*/
      console.log(searchReq)
      /*eslint-enable*/
      axios({
        method: "POST",
        url: "/api/v1/getNode",
        data: searchReq,
        headers: {"content-type": "text/plain"}
      }).then(result => {
        // this.response = result.data;
        /*eslint-disable*/
        var data = []
        console.log(result.data)
        var searchRsp = result.data["nodeInfo"]
        searchRsp.forEach(function(val) {
        console.log(val)
        data.push(
            {
              id: val["id"],
              name: val["name"],
              createTime: val["createTime"],
              updateTime: val["updateTime"],
              desc: val["desc"],
              type: val["type"],
              score: val["score"],
            } )
        } )
        this.data = data

        /*eslint-enable*/
      }).catch(error => {
        /*eslint-disable*/
        console.error(error);
        /*eslint-enable*/
      });

      console.log(this.data)
    },
    AddNode: function () {
      alert("开发中")
    },

    ShowMindNodes: function () {
      alert("开发中")
    }
  },
} );
</script>

<style>
#components-layout-demo-side .logo {
  height: 32px;
  margin: 16px;
  background: rgba(255, 255, 255, 0.3);
}

.site-layout .site-layout-background {
  background: #fff;
}
[data-theme='dark'] .site-layout .site-layout-background {
  background: #141414;
}
</style>
