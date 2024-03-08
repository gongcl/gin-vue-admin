<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Department字段:" prop="Department">
          <el-input v-model="formData.Department" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="YusuanType字段:" prop="YusuanType">
          <el-input v-model="formData.YusuanType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ProjectCode字段:" prop="ProjectCode">
          <el-input v-model="formData.ProjectCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ProjectName字段:" prop="ProjectName">
          <el-input v-model="formData.ProjectName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="CustomGroup字段:" prop="CustomGroup">
          <el-input v-model="formData.CustomGroup" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="CustomerName字段:" prop="CustomerName">
          <el-input v-model="formData.CustomerName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="TotalNR字段:" prop="TotalNR">
          <el-input-number v-model="formData.TotalNR" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="January字段:" prop="January">
          <el-input-number v-model="formData.January" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="February字段:" prop="February">
          <el-input-number v-model="formData.February" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="March字段:" prop="March">
          <el-input-number v-model="formData.March" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="April字段:" prop="April">
          <el-input-number v-model="formData.April" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="May字段:" prop="May">
          <el-input-number v-model="formData.May" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="June字段:" prop="June">
          <el-input-number v-model="formData.June" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="July字段:" prop="July">
          <el-input-number v-model="formData.July" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="August字段:" prop="August">
          <el-input-number v-model="formData.August" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="September字段:" prop="September">
          <el-input-number v-model="formData.September" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="October字段:" prop="October">
          <el-input-number v-model="formData.October" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="November字段:" prop="November">
          <el-input-number v-model="formData.November" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="December字段:" prop="December">
          <el-input-number v-model="formData.December" :precision="2" :clearable="true"></el-input-number>
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
  createMsYsXm24y,
  updateMsYsXm24y,
  findMsYsXm24y
} from '@/api/msYsXm24y'

defineOptions({
    name: 'MsYsXm24yForm'
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
            Department: '',
            YusuanType: '',
            ProjectCode: '',
            ProjectName: '',
            CustomGroup: '',
            CustomerName: '',
            TotalNR: 0,
            January: 0,
            February: 0,
            March: 0,
            April: 0,
            May: 0,
            June: 0,
            July: 0,
            August: 0,
            September: 0,
            October: 0,
            November: 0,
            December: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMsYsXm24y({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remsYsXm24y
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
               res = await createMsYsXm24y(formData.value)
               break
             case 'update':
               res = await updateMsYsXm24y(formData.value)
               break
             default:
               res = await createMsYsXm24y(formData.value)
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
