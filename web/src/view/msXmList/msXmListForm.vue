<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="客户名称:" prop="khmc">
          <el-input v-model="formData.khmc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="注册编码:" prop="zcbm">
          <el-input v-model="formData.zcbm" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合同编码:" prop="htbm">
          <el-input v-model="formData.htbm" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合同名称:" prop="htmc">
          <el-input v-model="formData.htmc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合同金额:" prop="htje">
          <el-input-number v-model="formData.htje" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="项目编码:" prop="xmbm">
          <el-input v-model="formData.xmbm" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="项目名称:" prop="xmmc">
          <el-input v-model="formData.xmmc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="项目金额:" prop="xmje">
          <el-input-number v-model="formData.xmje" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="项目开始日期:" prop="kssj">
          <el-date-picker v-model="formData.kssj" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="项目结束日期:" prop="jssj">
          <el-date-picker v-model="formData.jssj" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="项目经理:" prop="xmjl">
          <el-input v-model="formData.xmjl" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执行部门:" prop="zxbm">
          <el-input v-model="formData.zxbm" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="项目状态:" prop="xmzt">
          <el-input v-model="formData.xmzt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="所在城市:" prop="xmcs">
          <el-input v-model="formData.xmcs" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="在岸离岸:" prop="site">
          <el-input v-model="formData.site" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="结算模式:" prop="jsms">
          <el-input v-model="formData.jsms" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预算收入:" prop="srys">
          <el-input-number v-model="formData.srys" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="预算毛利:" prop="mlys">
          <el-input-number v-model="formData.mlys" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="预算毛利率:" prop="mlyys">
          <el-input-number v-model="formData.mlyys" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="滚动收入:" prop="gdsr">
          <el-input-number v-model="formData.gdsr" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="滚动毛利:" prop="gdml">
          <el-input-number v-model="formData.gdml" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="滚动毛利率:" prop="gdmll">
          <el-input-number v-model="formData.gdmll" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="滚动人工:" prop="gdrg">
          <el-input-number v-model="formData.gdrg" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="滚动其他成本:" prop="gdqt">
          <el-input-number v-model="formData.gdqt" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="状态标记:" prop="ztbj">
          <el-input v-model="formData.ztbj" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注说明:" prop="bzsm">
          <el-input v-model="formData.bzsm" :clearable="true" placeholder="请输入" />
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
  createMsXmList,
  updateMsXmList,
  findMsXmList
} from '@/api/msXmList'

defineOptions({
    name: 'MsXmListForm'
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
            khmc: '',
            zcbm: '',
            htbm: '',
            htmc: '',
            htje: 0,
            xmbm: '',
            xmmc: '',
            xmje: 0,
            kssj: new Date(),
            jssj: new Date(),
            xmjl: '',
            zxbm: '',
            xmzt: '',
            xmcs: '',
            site: '',
            jsms: '',
            srys: 0,
            mlys: 0,
            mlyys: 0,
            gdsr: 0,
            gdml: 0,
            gdmll: 0,
            gdrg: 0,
            gdqt: 0,
            ztbj: '',
            bzsm: '',
        })
// 验证规则
const rule = reactive({
               zcbm : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               xmbm : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               xmmc : [{
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
      const res = await findMsXmList({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remsXmList
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
               res = await createMsXmList(formData.value)
               break
             case 'update':
               res = await updateMsXmList(formData.value)
               break
             default:
               res = await createMsXmList(formData.value)
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
