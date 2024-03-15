<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="执行部门" prop="dept">
         <el-input v-model="searchInfo.dept" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="合同名称" prop="htmc">
         <el-input v-model="searchInfo.htmc" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="结束日期" prop="jsrq">
            
            <template #label>
            <span>
              结束日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startJsrq" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endJsrq ? time.getTime() > searchInfo.endJsrq.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endJsrq" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startJsrq ? time.getTime() < searchInfo.startJsrq.getTime() : false"></el-date-picker>

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>

          <!-- <el-button type="warning" @click="query(1,this)" circle >1月</el-button>
          <el-button type="warning" @click="query(2)" circle >2月</el-button>
          <el-button type="warning" @click="query(3)" circle >3月</el-button>
          <el-button type="warning" @click="query(4)" circle >4月</el-button>
          <el-button type="warning" @click="query(5)" circle >5月</el-button>
          <el-button type="warning" @click="query(6)" circle >6月</el-button>
          <el-button type="warning" @click="query(7)" circle >7月</el-button>
          <el-button type="warning" @click="query(8)" circle >8月</el-button>
          <el-button type="warning" @click="query(9)" circle >9月</el-button>
          <el-button type="warning" @click="query(10)" circle >10月</el-button>
          <el-button type="warning" @click="query(11)" circle >11月</el-button>
          <el-button type="warning" @click="query(12)" circle >12月</el-button> -->

          <el-button type="warning" :plain="isMonthSelected(month)" v-for="month in months" :key="month" @click="query(month)">{{month}}月</el-button>

        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list" style="display: none;">
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
        <!-- show-summary="true" -->
        <el-table
        ref="multipleTable"
        style="width: 100%;height:550px;"
        tooltip-effect="dark"
        stripe="true"
        border="true"
        fit="true"
        highlight-current-row="true"
        :data="tableData"
        show-overflow-tooltip="true"
        row-key="htmc"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <!-- <el-table-column type="selection" width="55" /> -->
        <el-table-column align="center" label="序号" type="index" width="55" /> 
        <!-- <el-table-column sortable align="left" label="执行部门" prop="dept" width="120" /> -->
        <el-table-column align="left" label="客户系" prop="kuxl" width="120" />
        <!-- <el-table-column sortable align="left" label="客户名称" prop="khmc" width="120" /> -->
        <!-- <el-table-column align="left" label="合同编码" prop="htbm" width="120" /> -->
        <el-table-column align="left" label="合同名称" prop="htmc" width="320" />
        <el-table-column align="left" label="合同属性" prop="htsx" width="120" />
        <el-table-column align="center" label="计费类型" prop="jflx" width="120" />
        <el-table-column align="center" label="币种" prop="jsbz" width="60" />
        <el-table-column align="center" label="合同金额" prop="xmje" width="120" />
        <el-table-column align="left" label="生效日期" width="110">
            <template #default="scope">{{ formatDate(scope.row.sxrq) }}</template>
         </el-table-column>
         <el-table-column sortable align="left" label="结束日期" width="110">
            <template #default="scope">{{ formatDate(scope.row.jsrq) }}</template>
         </el-table-column>
         <!-- <el-table-column align="left" label="结算周期" prop="jszq" width="80" />
        <el-table-column align="left" label="回款周期" prop="hkzq" width="60" /> -->
         <el-table-column align="left" label="合同拿回" width="110">
            <template #default="scope">{{ formatDate(scope.row.henh) }}</template>
         </el-table-column>
        <!-- <el-table-column align="left" label="是否B类" prop="sfbl" width="20" />
        <el-table-column align="left" label="是否E类" prop="sfel" width="20" /> -->
        <el-table-column align="left" label="备注说明" prop="bzsm" width="240" />
        <el-table-column align="left" label="操作" fixed="right" min-width="90">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateMsHtListFunc(scope.row)"></el-button>
            <!-- <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button> -->
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[100, 200]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="执行部门:"  prop="dept" >
              <el-input v-model="formData.dept" :clearable="true"  placeholder="请输入执行部门" />
            </el-form-item>
            <el-form-item label="客户系:"  prop="kuxl" >
              <el-input v-model="formData.kuxl" :clearable="true"  placeholder="请输入客户系" />
            </el-form-item>
            <el-form-item label="客户名称:"  prop="khmc" >
              <el-input v-model="formData.khmc" :clearable="true"  placeholder="请输入客户名称" />
            </el-form-item>
            <el-form-item label="合同编码:"  prop="htbm" >
              <el-input v-model="formData.htbm" :clearable="true"  placeholder="请输入合同编码" />
            </el-form-item>
            <el-form-item label="合同名称:"  prop="htmc" >
              <el-input v-model="formData.htmc" :clearable="true"  placeholder="请输入合同名称" />
            </el-form-item>
            <el-form-item label="回款周期:"  prop="hkzq" >
              <el-input v-model.number="formData.hkzq" :clearable="true" placeholder="请输入回款周期" />
            </el-form-item>
            <el-form-item label="合同属性:"  prop="htsx" >
              <el-input v-model="formData.htsx" :clearable="true"  placeholder="请输入合同属性" />
            </el-form-item>
            <el-form-item label="计费类型:"  prop="jflx" >
              <el-input v-model="formData.jflx" :clearable="true"  placeholder="请输入计费类型" />
            </el-form-item>
            <el-form-item label="结算币种:"  prop="jsbz" >
              <el-input v-model="formData.jsbz" :clearable="true"  placeholder="请输入结算币种" />
            </el-form-item>
            <el-form-item label="合同金额:"  prop="xmje" >
              <el-input-number v-model="formData.xmje"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="生效日期:"  prop="sxrq" >
              <el-date-picker v-model="formData.sxrq" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="结束日期:"  prop="jsrq" >
              <el-date-picker v-model="formData.jsrq" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="结算周期:"  prop="jszq" >
              <el-input v-model="formData.jszq" :clearable="true"  placeholder="请输入结算周期" />
            </el-form-item>
            <el-form-item label="合同拿回:"  prop="henh" >
              <el-date-picker v-model="formData.henh" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="是否B类:"  prop="sfbl" >
              <el-input v-model="formData.sfbl" :clearable="true"  placeholder="请输入是否B类" />
            </el-form-item>
            <el-form-item label="是否E类:"  prop="sfel" >
              <el-input v-model="formData.sfel" :clearable="true"  placeholder="请输入是否E类" />
            </el-form-item>
            <el-form-item label="备注说明:"  prop="bzsm" >
              <el-input v-model="formData.bzsm" :clearable="true"  placeholder="请输入跟踪情况" />
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
                <el-descriptions-item label="执行部门">
                        {{ formData.dept }}
                </el-descriptions-item>
                <el-descriptions-item label="客户系">
                        {{ formData.kuxl }}
                </el-descriptions-item>
                <el-descriptions-item label="客户名称">
                        {{ formData.khmc }}
                </el-descriptions-item>
                <el-descriptions-item label="合同编码">
                        {{ formData.htbm }}
                </el-descriptions-item>
                <el-descriptions-item label="合同名称">
                        {{ formData.htmc }}
                </el-descriptions-item>
                <el-descriptions-item label="合同属性">
                        {{ formData.htsx }}
                </el-descriptions-item>
                <el-descriptions-item label="计费类型">
                        {{ formData.jflx }}
                </el-descriptions-item>
                <el-descriptions-item label="结算币种">
                        {{ formData.jsbz }}
                </el-descriptions-item>
                <el-descriptions-item label="合同金额">
                        {{ formData.xmje }}
                </el-descriptions-item>
                <el-descriptions-item label="生效日期">
                      {{ formatDate(formData.sxrq) }}
                </el-descriptions-item>
                <el-descriptions-item label="结束日期">
                      {{ formatDate(formData.jsrq) }}
                </el-descriptions-item>
                <el-descriptions-item label="结算周期">
                        {{ formData.jszq }}
                </el-descriptions-item>
                <el-descriptions-item label="回款周期">
                        {{ formData.hkzq }}
                </el-descriptions-item>
                <el-descriptions-item label="合同拿回">
                      {{ formatDate(formData.henh) }}
                </el-descriptions-item>
                <el-descriptions-item label="是否B类">
                        {{ formData.sfbl }}
                </el-descriptions-item>
                <el-descriptions-item label="是否E类">
                        {{ formData.sfel }}
                </el-descriptions-item>
                <el-descriptions-item label="备注说明">
                        {{ formData.bzsm }}
                </el-descriptions-item>
                
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createMsHtList,
  deleteMsHtList,
  deleteMsHtListByIds,
  updateMsHtList,
  findMsHtList,
  getMsHtListList
} from '@/api/msHtList'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'MsHtList'
})


// searchInfo.value.startCreatedAt = '2024-01-01'
// searchInfo.value.endCreatedAt = '2024-07-01'


// 自动化生成的字典（可能为空）以及字段
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
        bzsm:''
        })


// 验证规则
const rule = reactive({
               htmc : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
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
        jsrq : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startJsrq && !searchInfo.value.endJsrq) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startJsrq && searchInfo.value.endJsrq) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startJsrq && searchInfo.value.endJsrq && (searchInfo.value.startJsrq.getTime() === searchInfo.value.endJsrq.getTime() || searchInfo.value.startJsrq.getTime() > searchInfo.value.endJsrq.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(100)
const tableData = ref([])
const searchInfo = ref({})

const months = [1,2,3,4,5,6,7,8,9,10,11,12]
//["1月","2月","3月","4月","5月","6月","7月","8月","9月","10月","11月","12月"]

// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            dept: 'dept',
            jsrq: 'jsrq',
            khmc: 'khmc',
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

let selmonth = 0

const query = (i, item) =>{

  selmonth = i
  
  searchInfo.value.startJsrq = new Date('2024-'+i+'-01')
  searchInfo.value.endJsrq = new Date('2024-'+i+'-31')
  getTableData()
  
}

const isMonthSelected = (month)=>{
console.log(selmonth)
  return !(month == selmonth)
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 100
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
  const table = await getMsHtListList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteMsHtListFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const htmcs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          htmcs.push(item.htmc)
        })
      const res = await deleteMsHtListByIds({ htmcs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === htmcs.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMsHtListFunc = async(row) => {
    const res = await findMsHtList({ htmc: row.htmc })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.remsHtList
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMsHtListFunc = async (row) => {
    const res = await deleteMsHtList({ htmc: row.htmc })
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
  const res = await findMsHtList({ htmc: row.htmc })
  if (res.code === 0) {
    formData.value = res.data.remsHtList
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
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
          bzsm:''
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
        bzsm:''
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
