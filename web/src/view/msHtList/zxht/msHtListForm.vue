<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="执行部门:" prop="dept">
          <el-input v-model="formData.dept" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合同编码:" prop="htbm">
          <el-input v-model="formData.htbm" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合同名称:" prop="htmc">
          <el-input v-model="formData.htmc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="回款周期:" prop="hkzq">
          <el-input v-model.number="formData.hkzq" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合同拿回:" prop="henh">
          <el-date-picker v-model="formData.henh" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="合同属性:" prop="htsx">
          <el-input v-model="formData.htsx" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="计费类型:" prop="jflx">
          <el-input v-model="formData.jflx" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="结算币种:" prop="jsbz">
          <el-input v-model="formData.jsbz" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="结束日期:" prop="jsrq">
          <el-date-picker v-model="formData.jsrq" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="结算周期:" prop="jszq">
          <el-input v-model="formData.jszq" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客户名称:" prop="khmc">
          <el-input v-model="formData.khmc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客户系:" prop="kuxl">
          <el-input v-model="formData.kuxl" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否B类:" prop="sfbl">
          <el-input v-model="formData.sfbl" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否E类:" prop="sfel">
          <el-input v-model="formData.sfel" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="生效日期:" prop="sxrq">
          <el-date-picker v-model="formData.sxrq" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="合同金额:" prop="xmje">
          <el-input-number v-model="formData.xmje" :precision="2" :clearable="true"></el-input-number>
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
  createMsHtList,
  updateMsHtList,
  findMsHtList
} from '@/api/msHtList'

defineOptions({
    name: 'MsHtListForm'
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
            dept: '',
            htbm: '',
            htmc: '',
            hkzq: 0,
            henh: new Date(),
            htsx: '',
            jflx: '',
            jsbz: '',
            jsrq: new Date(),
            jszq: '',
            khmc: '',
            kuxl: '',
            sfbl: '',
            sfel: '',
            sxrq: new Date(),
            xmje: 0,
        })
// 验证规则
const rule = reactive({
               htmc : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMsHtList({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remsHtList
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
               res = await createMsHtList(formData.value)
               break
             case 'update':
               res = await updateMsHtList(formData.value)
               break
             default:
               res = await createMsHtList(formData.value)
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
