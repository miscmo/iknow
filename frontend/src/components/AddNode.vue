<template>
  <div>
    <a-button type="primary" @click="showModal">添加</a-button>
    <a-modal v-model:visible="visible" width="1000px" title="Title" @ok="ensureAddNode">
      <template #footer>
        <a-button key="back" @click="cancelAddNode">取消</a-button>
        <a-button key="submit" type="primary" :loading="loading" @click="ensureAddNode">添加</a-button>
      </template>
      <a-form
          ref="formRef"
          :model="formState"
          :rules="rules"
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
      >
        <a-form-item ref="name" label="名称" name="name">
          <a-input v-model:value="formState.name" />
        </a-form-item>
        <a-form-item label="类型" name="type">
          <a-select v-model:value="formState.type" placeholder="请选择：">
            <a-select-option :value='0'>默认</a-select-option>
            <a-select-option :value='1'>笔记</a-select-option>
            <a-select-option :value='2'>计划</a-select-option>
            <a-select-option :value='3'>练习</a-select-option>
          </a-select>
        </a-form-item>
<!--        <a-form-item label="Activity time" required name="date1">-->
<!--          <a-date-picker-->
<!--              v-model:value="formState.date1"-->
<!--              show-time-->
<!--              type="date"-->
<!--              placeholder="Pick a date"-->
<!--              style="width: 100%"-->
<!--          />-->
<!--        </a-form-item>-->
        <a-form-item label="折叠" name="collapse">
          <a-switch v-model:checked="formState.collapse" />
        </a-form-item>

<!--        <a-form-item label="权限" name="permission">-->
<!--          <a-checkbox-group v-model:value="formState.permission">-->
<!--            <a-checkbox value="1" name="permission">无限制</a-checkbox>-->
<!--            <a-checkbox value="2" name="permission">仅自己可写</a-checkbox>-->
<!--            <a-checkbox value="3" name="permission">仅自己可读</a-checkbox>-->
<!--          </a-checkbox-group>-->

<!--        </a-form-item>-->
        <a-form-item label="权限" name="permission">
          <a-radio-group v-model:value="formState.permission">
            <a-radio :value='0'>无限制</a-radio>
            <a-radio :value='1'>仅自己可写</a-radio>
            <a-radio :value='2'>仅自己可读</a-radio>
          </a-radio-group>
        </a-form-item>

        <a-form-item label="描述" name="desc">
          <a-textarea v-model:value="formState.desc" />
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
          <a-button type="primary" @click="onSubmit">Create</a-button>
          <a-button style="margin-left: 10px" @click="resetForm">Reset</a-button>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script>
import { defineComponent, reactive, ref, toRaw } from 'vue';
import {AddNodes} from "@/api/Nodes";
export default defineComponent({
  name: "AddNodes",
  setup() {
    const loading = ref(false);
    const visible = ref(false);

    const showModal = () => {
      visible.value = true;
    };

    const ensureAddNode = () => {
      loading.value = true;
      setTimeout(() => {
        loading.value = false;
        visible.value = false;
      }, 2000);
    };

    const cancelAddNode = () => {
      visible.value = false;
    };

    const formRef = ref();
    const formState = reactive({
      name: '',
      type: 0,
      notes: [],
      subNodes: [],
      collapse: false,
      score: 0,
      permission: 0,
      desc: '',
      createTime: '',
      updateTime: '',
    });
    const rules = {
      name: [ { required: true, message: '请输入节点名称', trigger: 'blur', },
        { min: 3, /*max: 5,*/ message: '长度必须大于3', trigger: 'blur', }, ],
      //type: [ { required: true, message: '请选择节点类型', trigger: 'change', }, ],
      //permissions: [ { required: true, message: '请选择节点类型', trigger: 'change', }, ],
      // subNotes: [ { required: true, message: 'Please pick a date', trigger: 'change', type: 'object', }, ],
      // type: [ { type: 'array', required: true, message: 'Please select at least one activity type', trigger: 'change', }, ],
      // resource: [ { required: true, message: 'Please select activity resource', trigger: 'change', }, ],
      desc: [ { required: false, message: '', trigger: 'blur', }, ],
    };

    const onSubmit = () => {
      formRef.value
          .validate()
          .then(() => {
            console.log('values', formState, toRaw(formState));
            var req = {"node": formState}
            console.log(req)

            AddNodes(req).then(result => {
              console.log(result.data)
            }).catch(error => {
              console.error(error)
            })

          })
          .catch(error => {
            console.log('error', error);
          });
    };

    const resetForm = () => {
      formRef.value.resetFields();
    };

    return {
      formRef,
      labelCol: {
        span: 4,
      },
      wrapperCol: {
        span: 14,
      },
      other: '',
      formState,
      rules,
      onSubmit,
      resetForm,
      loading,
      visible,
      showModal,
      ensureAddNode,
      cancelAddNode,
    };

  },
});
</script>