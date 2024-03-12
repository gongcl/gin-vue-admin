<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="客户名称" prop="customerName">
         <el-input v-model="searchInfo.customerName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="交付部" prop="department">
         <el-input v-model="searchInfo.department" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="年度预算" prop="nrTotal">
            
             <el-input v-model.number="searchInfo.nrTotal" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="预算类型" prop="nrType">
         <el-input v-model="searchInfo.nrType" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="项目名称" prop="projectName">
         <el-input v-model="searchInfo.projectName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
<!--        <div class="gva-btn-list">-->
<!--            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>-->
<!--            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">-->
<!--            <p>确定要删除吗？</p>-->
<!--            <div style="text-align: right; margin-top: 8px;">-->
<!--                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>-->
<!--                <el-button type="primary" @click="onDelete">确定</el-button>-->
<!--            </div>-->
<!--            <template #reference>-->
<!--                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>-->
<!--            </template>-->
<!--            </el-popover>-->
<!--        </div>-->
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="projectName"
        @selection-change="handleSelectionChange"
        >
        <el-table-column align="center" label="序号" type="index" width="55" />
          <el-table-column align="left" label="交付部" prop="department" width="120" />
        <el-table-column align="left" label="客户系" prop="customerGroup" width="120" />
        <el-table-column align="left" label="客户名称" disabled="true" prop="customerName" width="120" />
         <el-table-column align="left" label="项目编码" v-if="false" prop="projectCode" width="120" />
         <el-table-column align="left" label="项目名称" prop="projectName" width="120" />
          <el-table-column align="left" label="预算类型" prop="nrType" width="120" />
          <el-table-column align="left" label="年度预算" prop="nrTotal" width="120" />
        <el-table-column align="left" label="1月" prop="month1" width="120" />
        <el-table-column align="left" label="2月" prop="month2" width="120" />
        <el-table-column align="left" label="3月" prop="month3" width="120" />
        <el-table-column align="left" label="4月" prop="month4" width="120" />
        <el-table-column align="left" label="5月" prop="month5" width="120" />
        <el-table-column align="left" label="6月" prop="month6" width="120" />
        <el-table-column align="left" label="7月" prop="month7" width="120" />
        <el-table-column align="left" label="8月" prop="month8" width="120" />
        <el-table-column align="left" label="9月" prop="month9" width="120" />
        <el-table-column align="left" label="10月" prop="month10" width="120" />
        <el-table-column align="left" label="11月" prop="month11" width="120" />
        <el-table-column align="left" label="12月" prop="month12" width="120" />

<!--        <el-table-column align="left" label="操作" fixed="right" min-width="240">-->
<!--            <template #default="scope">-->
<!--            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">-->
<!--                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>-->
<!--                查看详情-->
<!--            </el-button>-->
<!--            <el-button type="primary" link icon="edit" class="table-button" @click="updateMsYsnrXmFunc(scope.row)">变更</el-button>-->
<!--            </template>-->
<!--        </el-table-column>-->
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="customerGroup字段:"  prop="customerGroup" >
              <el-input v-model="formData.customerGroup" :clearable="true"  placeholder="请输入customerGroup字段" />
            </el-form-item>
            <el-form-item label="customerName字段:"  prop="customerName" >
              <el-input v-model="formData.customerName" :clearable="true"  placeholder="请输入customerName字段" />
            </el-form-item>
            <el-form-item label="department字段:"  prop="department" >
              <el-input v-model="formData.department" :clearable="true"  placeholder="请输入department字段" />
            </el-form-item>
            <el-form-item label="month1字段:"  prop="month1" >
              <el-input-number v-model="formData.month1"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="month10字段:"  prop="month10" >
              <el-input-number v-model="formData.month10"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="month11字段:"  prop="month11" >
              <el-input-number v-model="formData.month11"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="month12字段:"  prop="month12" >
              <el-input-number v-model="formData.month12"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="month2字段:"  prop="month2" >
              <el-input-number v-model="formData.month2"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="month3字段:"  prop="month3" >
              <el-input-number v-model="formData.month3"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="month4字段:"  prop="month4" >
              <el-input-number v-model="formData.month4"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="month5字段:"  prop="month5" >
              <el-input-number v-model="formData.month5"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="month6字段:"  prop="month6" >
              <el-input-number v-model="formData.month6"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="month7字段:"  prop="month7" >
              <el-input-number v-model="formData.month7"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="month8字段:"  prop="month8" >
              <el-input-number v-model="formData.month8"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="month9字段:"  prop="month9" >
              <el-input-number v-model="formData.month9"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="nrTotal字段:"  prop="nrTotal" >
              <el-input-number v-model="formData.nrTotal"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="nrType字段:"  prop="nrType" >
              <el-input v-model="formData.nrType" :clearable="true"  placeholder="请输入nrType字段" />
            </el-form-item>
            <el-form-item label="projectCode字段:"  prop="projectCode" >
              <el-input v-model="formData.projectCode" :clearable="true"  placeholder="请输入projectCode字段" />
            </el-form-item>
            <el-form-item label="projectName字段:"  prop="projectName" >
              <el-input v-model="formData.projectName" :clearable="true"  placeholder="请输入projectName字段" />
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions :column="1" border>
                <el-descriptions-item label="customerGroup字段">
                        {{ formData.customerGroup }}
                </el-descriptions-item>
                <el-descriptions-item label="customerName字段">
                        {{ formData.customerName }}
                </el-descriptions-item>
                <el-descriptions-item label="department字段">
                        {{ formData.department }}
                </el-descriptions-item>
                <el-descriptions-item label="month1字段">
                        {{ formData.month1 }}
                </el-descriptions-item>
                <el-descriptions-item label="month10字段">
                        {{ formData.month10 }}
                </el-descriptions-item>
                <el-descriptions-item label="month11字段">
                        {{ formData.month11 }}
                </el-descriptions-item>
                <el-descriptions-item label="month12字段">
                        {{ formData.month12 }}
                </el-descriptions-item>
                <el-descriptions-item label="month2字段">
                        {{ formData.month2 }}
                </el-descriptions-item>
                <el-descriptions-item label="month3字段">
                        {{ formData.month3 }}
                </el-descriptions-item>
                <el-descriptions-item label="month4字段">
                        {{ formData.month4 }}
                </el-descriptions-item>
                <el-descriptions-item label="month5字段">
                        {{ formData.month5 }}
                </el-descriptions-item>
                <el-descriptions-item label="month6字段">
                        {{ formData.month6 }}
                </el-descriptions-item>
                <el-descriptions-item label="month7字段">
                        {{ formData.month7 }}
                </el-descriptions-item>
                <el-descriptions-item label="month8字段">
                        {{ formData.month8 }}
                </el-descriptions-item>
                <el-descriptions-item label="month9字段">
                        {{ formData.month9 }}
                </el-descriptions-item>
                <el-descriptions-item label="nrTotal字段">
                        {{ formData.nrTotal }}
                </el-descriptions-item>
                <el-descriptions-item label="nrType字段">
                        {{ formData.nrType }}
                </el-descriptions-item>
                <el-descriptions-item label="projectCode字段">
                        {{ formData.projectCode }}
                </el-descriptions-item>
                <el-descriptions-item label="projectName字段">
                        {{ formData.projectName }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createMsYsnrXm,
  deleteMsYsnrXm,
  deleteMsYsnrXmByIds,
  updateMsYsnrXm,
  findMsYsnrXm,
  getMsYsnrXmList
} from '@/api/msYsnrXm'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'MsYsnrXm'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        customerGroup: '',
        customerName: '',
        department: '',
        month1: 0,
        month10: 0,
        month11: 0,
        month12: 0,
        month2: 0,
        month3: 0,
        month4: 0,
        month5: 0,
        month6: 0,
        month7: 0,
        month8: 0,
        month9: 0,
        nrTotal: 0,
        nrType: '',
        projectCode: '',
        projectName: '',
        })


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getMsYsnrXmList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteMsYsnrXmFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const projectNames = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          projectNames.push(item.projectName)
        })
      const res = await deleteMsYsnrXmByIds({ projectNames })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === projectNames.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMsYsnrXmFunc = async(row) => {
    const res = await findMsYsnrXm({ projectName: row.projectName })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.remsYsnrXm
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMsYsnrXmFunc = async (row) => {
    const res = await deleteMsYsnrXm({ projectName: row.projectName })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findMsYsnrXm({ projectName: row.projectName })
  if (res.code === 0) {
    formData.value = res.data.remsYsnrXm
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          customerGroup: '',
          customerName: '',
          department: '',
          month1: 0,
          month10: 0,
          month11: 0,
          month12: 0,
          month2: 0,
          month3: 0,
          month4: 0,
          month5: 0,
          month6: 0,
          month7: 0,
          month8: 0,
          month9: 0,
          nrTotal: 0,
          nrType: '',
          projectCode: '',
          projectName: '',
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        customerGroup: '',
        customerName: '',
        department: '',
        month1: 0,
        month10: 0,
        month11: 0,
        month12: 0,
        month2: 0,
        month3: 0,
        month4: 0,
        month5: 0,
        month6: 0,
        month7: 0,
        month8: 0,
        month9: 0,
        nrTotal: 0,
        nrType: '',
        projectCode: '',
        projectName: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
