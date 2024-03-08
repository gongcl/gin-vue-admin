<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="Department字段" prop="Department">
         <el-input v-model="searchInfo.Department" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="YusuanType字段" prop="YusuanType">
         <el-input v-model="searchInfo.YusuanType" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="ProjectName字段" prop="ProjectName">
         <el-input v-model="searchInfo.ProjectName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="CustomGroup字段" prop="CustomGroup">
         <el-input v-model="searchInfo.CustomGroup" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="CustomerName字段" prop="CustomerName">
         <el-input v-model="searchInfo.CustomerName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="TotalNR字段" prop="TotalNR">
            
             <el-input v-model.number="searchInfo.TotalNR" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="交付部" prop="Department" width="120" />
        <el-table-column align="left" label="预算类型" prop="YusuanType" width="120" />
        <el-table-column sortable align="left" label="项目编码" prop="ProjectCode" width="120" />
        <el-table-column sortable align="left" label="项目名称" prop="ProjectName" width="120" />
        <el-table-column align="left" label="客户系" prop="CustomGroup" width="120" />
        <el-table-column sortable align="left" label="客户名称" prop="CustomerName" width="120" />
        <el-table-column align="left" label="预算NR" prop="TotalNR" width="120" />
        <el-table-column align="left" label="1月" prop="January" width="120" />
        <el-table-column align="left" label="2月" prop="February" width="120" />
        <el-table-column align="left" label="3月" prop="March" width="120" />
        <el-table-column align="left" label="4月" prop="April" width="120" />
        <el-table-column align="left" label="5月" prop="May" width="120" />
        <el-table-column align="left" label="6月" prop="June" width="120" />
        <el-table-column align="left" label="7月" prop="July" width="120" />
        <el-table-column align="left" label="8月" prop="August" width="120" />
        <el-table-column align="left" label="9月" prop="September" width="120" />
        <el-table-column align="left" label="10月" prop="October" width="120" />
        <el-table-column align="left" label="11月" prop="November" width="120" />
        <el-table-column align="left" label="12月" prop="December" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateMsYsXm24yFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
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
            <el-form-item label="Department字段:"  prop="Department" >
              <el-input v-model="formData.Department" :clearable="true"  placeholder="请输入Department字段" />
            </el-form-item>
            <el-form-item label="YusuanType字段:"  prop="YusuanType" >
              <el-input v-model="formData.YusuanType" :clearable="true"  placeholder="请输入YusuanType字段" />
            </el-form-item>
            <el-form-item label="ProjectCode字段:"  prop="ProjectCode" >
              <el-input v-model="formData.ProjectCode" :clearable="true"  placeholder="请输入ProjectCode字段" />
            </el-form-item>
            <el-form-item label="ProjectName字段:"  prop="ProjectName" >
              <el-input v-model="formData.ProjectName" :clearable="true"  placeholder="请输入ProjectName字段" />
            </el-form-item>
            <el-form-item label="CustomGroup字段:"  prop="CustomGroup" >
              <el-input v-model="formData.CustomGroup" :clearable="true"  placeholder="请输入CustomGroup字段" />
            </el-form-item>
            <el-form-item label="CustomerName字段:"  prop="CustomerName" >
              <el-input v-model="formData.CustomerName" :clearable="true"  placeholder="请输入CustomerName字段" />
            </el-form-item>
            <el-form-item label="TotalNR字段:"  prop="TotalNR" >
              <el-input-number v-model="formData.TotalNR"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="January字段:"  prop="January" >
              <el-input-number v-model="formData.January"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="February字段:"  prop="February" >
              <el-input-number v-model="formData.February"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="March字段:"  prop="March" >
              <el-input-number v-model="formData.March"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="April字段:"  prop="April" >
              <el-input-number v-model="formData.April"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="May字段:"  prop="May" >
              <el-input-number v-model="formData.May"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="June字段:"  prop="June" >
              <el-input-number v-model="formData.June"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="July字段:"  prop="July" >
              <el-input-number v-model="formData.July"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="August字段:"  prop="August" >
              <el-input-number v-model="formData.August"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="September字段:"  prop="September" >
              <el-input-number v-model="formData.September"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="October字段:"  prop="October" >
              <el-input-number v-model="formData.October"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="November字段:"  prop="November" >
              <el-input-number v-model="formData.November"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="December字段:"  prop="December" >
              <el-input-number v-model="formData.December"  style="width:100%" :precision="2" :clearable="true"  />
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
                <el-descriptions-item label="Department字段">
                        {{ formData.Department }}
                </el-descriptions-item>
                <el-descriptions-item label="YusuanType字段">
                        {{ formData.YusuanType }}
                </el-descriptions-item>
                <el-descriptions-item label="ProjectCode字段">
                        {{ formData.ProjectCode }}
                </el-descriptions-item>
                <el-descriptions-item label="ProjectName字段">
                        {{ formData.ProjectName }}
                </el-descriptions-item>
                <el-descriptions-item label="CustomGroup字段">
                        {{ formData.CustomGroup }}
                </el-descriptions-item>
                <el-descriptions-item label="CustomerName字段">
                        {{ formData.CustomerName }}
                </el-descriptions-item>
                <el-descriptions-item label="TotalNR字段">
                        {{ formData.TotalNR }}
                </el-descriptions-item>
                <el-descriptions-item label="1月">
                        {{ formData.January }}
                </el-descriptions-item>
                <el-descriptions-item label="2月">
                        {{ formData.February }}
                </el-descriptions-item>
                <el-descriptions-item label="3月">
                        {{ formData.March }}
                </el-descriptions-item>
                <el-descriptions-item label="4月">
                        {{ formData.April }}
                </el-descriptions-item>
                <el-descriptions-item label="5月">
                        {{ formData.May }}
                </el-descriptions-item>
                <el-descriptions-item label="6月">
                        {{ formData.June }}
                </el-descriptions-item>
                <el-descriptions-item label="July字段">
                        {{ formData.July }}
                </el-descriptions-item>
                <el-descriptions-item label="August字段">
                        {{ formData.August }}
                </el-descriptions-item>
                <el-descriptions-item label="September字段">
                        {{ formData.September }}
                </el-descriptions-item>
                <el-descriptions-item label="October字段">
                        {{ formData.October }}
                </el-descriptions-item>
                <el-descriptions-item label="November字段">
                        {{ formData.November }}
                </el-descriptions-item>
                <el-descriptions-item label="December字段">
                        {{ formData.December }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createMsYsXm24y,
  deleteMsYsXm24y,
  deleteMsYsXm24yByIds,
  updateMsYsXm24y,
  findMsYsXm24y,
  getMsYsXm24yList
} from '@/api/msYsXm24y'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'MsYsXm24y'
})

// 自动化生成的字典（可能为空）以及字段
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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            ProjectCode: 'Project_code',
            ProjectName: 'Project_name',
            CustomerName: 'Customer_name',
  }

  let sort = sortMap[prop]
  if(!sort){
  //  sort = prop.replace(/[A-Z]/g, match => _${match.toLowerCase()})
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

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
  const table = await getMsYsXm24yList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteMsYsXm24yFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteMsYsXm24yByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMsYsXm24yFunc = async(row) => {
    const res = await findMsYsXm24y({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.remsYsXm24y
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMsYsXm24yFunc = async (row) => {
    const res = await deleteMsYsXm24y({ ID: row.ID })
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
  const res = await findMsYsXm24y({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.remsYsXm24y
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
