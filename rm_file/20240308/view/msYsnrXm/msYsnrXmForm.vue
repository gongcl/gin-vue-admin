<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="department字段:" prop="department">
          <el-input v-model="formData.department" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="nrType字段:" prop="nrType">
          <el-input v-model="formData.nrType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="projectCode字段:" prop="projectCode">
          <el-input v-model="formData.projectCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="projectName字段:" prop="projectName">
          <el-input v-model="formData.projectName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="customerGroup字段:" prop="customerGroup">
          <el-input v-model="formData.customerGroup" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="customerName字段:" prop="customerName">
          <el-input v-model="formData.customerName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="nrTotal字段:" prop="nrTotal">
          <el-input-number v-model="formData.nrTotal" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month1字段:" prop="month1">
          <el-input-number v-model="formData.month1" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month2字段:" prop="month2">
          <el-input-number v-model="formData.month2" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month3字段:" prop="month3">
          <el-input-number v-model="formData.month3" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month4字段:" prop="month4">
          <el-input-number v-model="formData.month4" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month5字段:" prop="month5">
          <el-input-number v-model="formData.month5" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month6字段:" prop="month6">
          <el-input-number v-model="formData.month6" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month7字段:" prop="month7">
          <el-input-number v-model="formData.month7" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month8字段:" prop="month8">
          <el-input-number v-model="formData.month8" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month9字段:" prop="month9">
          <el-input-number v-model="formData.month9" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month10字段:" prop="month10">
          <el-input-number v-model="formData.month10" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month11字段:" prop="month11">
          <el-input-number v-model="formData.month11" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="month12字段:" prop="month12">
          <el-input-number v-model="formData.month12" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createMsYsnrXm,
  updateMsYsnrXm,
  findMsYsnrXm
} from '@/api/msYsnrXm'

defineOptions({
    name: 'MsYsnrXmForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            department: '',
            nrType: '',
            projectCode: '',
            projectName: '',
            customerGroup: '',
            customerName: '',
            nrTotal: 0,
            month1: 0,
            month2: 0,
            month3: 0,
            month4: 0,
            month5: 0,
            month6: 0,
            month7: 0,
            month8: 0,
            month9: 0,
            month10: 0,
            month11: 0,
            month12: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMsYsnrXm({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remsYsnrXm
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createMsYsnrXm(formData.value)
               break
             case 'update':
               res = await updateMsYsnrXm(formData.value)
               break
             default:
               res = await createMsYsnrXm(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
