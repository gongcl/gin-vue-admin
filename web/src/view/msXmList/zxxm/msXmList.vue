<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      
      
        <el-form-item label="客户名称" prop="khmc">
         <el-input v-model="searchInfo.khmc" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="项目名称" prop="xmmc">
         <el-input v-model="searchInfo.xmmc" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="项目结束日期" prop="jssj">
            
            <template #label>
            <span>
              项目结束日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startJssj" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endJssj ? time.getTime() > searchInfo.endJssj.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endJssj" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startJssj ? time.getTime() < searchInfo.startJssj.getTime() : false"></el-date-picker>

        </el-form-item>
        <el-form-item label="项目经理" prop="xmjl">
         <el-input v-model="searchInfo.xmjl" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="执行部门" prop="zxbm">
         <el-input v-model="searchInfo.zxbm" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="所在城市" prop="xmcs">
         <el-input v-model="searchInfo.xmcs" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="滚动毛利率" prop="gdmll">
            
             <el-input v-model.number="searchInfo.gdmll" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="状态标记" prop="ztbj">
         <el-input v-model="searchInfo.ztbj" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <!-- <div class="gva-btn-list">
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
        </div> -->
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
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column align="center" label="序号" type="index" width="55" />
        
        <el-table-column sortable align="left" label="客户名称" prop="khmc" width="140" />
        <!-- <el-table-column align="left" label="注册编码" prop="zcbm" width="120" />
        <el-table-column align="left" label="合同编码" prop="htbm" width="120" /> -->
        <!-- <el-table-column align="left" label="合同名称" prop="htmc" width="120" /> -->
        <!-- <el-table-column align="left" label="合同金额" prop="htje" width="120" /> -->
        <!-- <el-table-column align="left" label="项目编码" prop="xmbm" width="120" /> -->
        <el-table-column align="left" label="项目名称" prop="xmmc" width="220" />
        <el-table-column align="center" label="项目金额" prop="xmje" width="100" />
         <el-table-column align="left" label="开始日期" width="100">
            <template #default="scope">{{ formatDate(scope.row.kssj) }}</template>
         </el-table-column>
         <el-table-column align="left" label="结束日期" width="100">
            <template #default="scope">{{ formatDate(scope.row.jssj) }}</template>
         </el-table-column>
        <el-table-column align="left" label="项目经理" prop="xmjl" width="120" />
        <!-- <el-table-column align="left" label="执行部门" prop="zxbm" width="120" /> -->
        <el-table-column align="left" label="项目状态" prop="xmzt" width="100" />
        <!-- <el-table-column align="left" label="所在城市" prop="xmcs" width="120" />
        <el-table-column align="left" label="在岸离岸" prop="site" width="120" /> -->
        <el-table-column align="left" label="结算模式" prop="jsms" width="100" />
        <el-table-column align="center" label="预算收入" prop="srys" width="90" />
        <el-table-column align="center" label="预算毛利" prop="mlys" width="90" />
        <!-- <el-table-column align="left" label="预算毛利率" prop="mlyys" width="120" /> -->
        <el-table-column align="center" label="滚动NR" prop="gdsr" width="90" />
        <el-table-column align="center" label="滚动NP" prop="gdml" width="90" />
        <el-table-column align="center" label="滚动NP%" prop="gdmll" width="90" />
        <!-- <el-table-column align="left" label="滚动人工" prop="gdrg" width="120" />
        <el-table-column align="left" label="滚动其他成本" prop="gdqt" width="120" /> -->
        <el-table-column align="left" label="状态标记" prop="ztbj" width="120" />
        <el-table-column align="left" label="备注说明" prop="bzsm" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="90">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateMsXmListFunc(scope.row)"></el-button>
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
            <el-form-item label="客户名称:"  prop="khmc" >
              <el-input v-model="formData.khmc" :clearable="true"  placeholder="请输入客户名称" />
            </el-form-item>
            <el-form-item label="注册编码:"  prop="zcbm" >
              <el-input v-model="formData.zcbm" :clearable="true"  placeholder="请输入注册编码" />
            </el-form-item>
            <el-form-item label="合同编码:"  prop="htbm" >
              <el-input v-model="formData.htbm" :clearable="true"  placeholder="请输入合同编码" />
            </el-form-item>
            <el-form-item label="合同名称:"  prop="htmc" >
              <el-input v-model="formData.htmc" :clearable="true"  placeholder="请输入合同名称" />
            </el-form-item>
            <el-form-item label="合同金额:"  prop="htje" >
              <el-input-number v-model="formData.htje"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="项目编码:"  prop="xmbm" >
              <el-input v-model="formData.xmbm" :clearable="true"  placeholder="请输入项目编码" />
            </el-form-item>
            <el-form-item label="项目名称:"  prop="xmmc" >
              <el-input v-model="formData.xmmc" :clearable="true"  placeholder="请输入项目名称" />
            </el-form-item>
            <el-form-item label="项目金额:"  prop="xmje" >
              <el-input-number v-model="formData.xmje"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="项目开始日期:"  prop="kssj" >
              <el-date-picker v-model="formData.kssj" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="项目结束日期:"  prop="jssj" >
              <el-date-picker v-model="formData.jssj" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="项目经理:"  prop="xmjl" >
              <el-input v-model="formData.xmjl" :clearable="true"  placeholder="请输入项目经理" />
            </el-form-item>
            <el-form-item label="执行部门:"  prop="zxbm" >
              <el-input v-model="formData.zxbm" :clearable="true"  placeholder="请输入执行部门" />
            </el-form-item>
            <el-form-item label="项目状态:"  prop="xmzt" >
              <el-input v-model="formData.xmzt" :clearable="true"  placeholder="请输入项目状态" />
            </el-form-item>
            <el-form-item label="所在城市:"  prop="xmcs" >
              <el-input v-model="formData.xmcs" :clearable="true"  placeholder="请输入所在城市" />
            </el-form-item>
            <el-form-item label="在岸离岸:"  prop="site" >
              <el-input v-model="formData.site" :clearable="true"  placeholder="请输入在岸离岸" />
            </el-form-item>
            <el-form-item label="结算模式:"  prop="jsms" >
              <el-input v-model="formData.jsms" :clearable="true"  placeholder="请输入结算模式" />
            </el-form-item>
            <el-form-item label="预算收入:"  prop="srys" >
              <el-input-number v-model="formData.srys"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="预算毛利:"  prop="mlys" >
              <el-input-number v-model="formData.mlys"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="预算毛利率:"  prop="mlyys" >
              <el-input-number v-model="formData.mlyys"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="滚动收入:"  prop="gdsr" >
              <el-input-number v-model="formData.gdsr"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="滚动毛利:"  prop="gdml" >
              <el-input-number v-model="formData.gdml"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="滚动毛利率:"  prop="gdmll" >
              <el-input-number v-model="formData.gdmll"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="滚动人工:"  prop="gdrg" >
              <el-input-number v-model="formData.gdrg"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="滚动其他成本:"  prop="gdqt" >
              <el-input-number v-model="formData.gdqt"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="状态标记:"  prop="ztbj" >
              <el-input v-model="formData.ztbj" :clearable="true"  placeholder="请输入状态标记" />
            </el-form-item>
            <el-form-item label="备注说明:"  prop="bzsm" >
              <el-input v-model="formData.bzsm" :clearable="true"  placeholder="请输入备注说明" />
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
                <el-descriptions-item label="客户名称">
                        {{ formData.khmc }}
                </el-descriptions-item>
                <el-descriptions-item label="注册编码">
                        {{ formData.zcbm }}
                </el-descriptions-item>
                <el-descriptions-item label="合同编码">
                        {{ formData.htbm }}
                </el-descriptions-item>
                <el-descriptions-item label="合同名称">
                        {{ formData.htmc }}
                </el-descriptions-item>
                <el-descriptions-item label="合同金额">
                        {{ formData.htje }}
                </el-descriptions-item>
                <el-descriptions-item label="项目编码">
                        {{ formData.xmbm }}
                </el-descriptions-item>
                <el-descriptions-item label="项目名称">
                        {{ formData.xmmc }}
                </el-descriptions-item>
                <el-descriptions-item label="项目金额">
                        {{ formData.xmje }}
                </el-descriptions-item>
                <el-descriptions-item label="项目开始日期">
                      {{ formatDate(formData.kssj) }}
                </el-descriptions-item>
                <el-descriptions-item label="项目结束日期">
                      {{ formatDate(formData.jssj) }}
                </el-descriptions-item>
                <el-descriptions-item label="项目经理">
                        {{ formData.xmjl }}
                </el-descriptions-item>
                <el-descriptions-item label="执行部门">
                        {{ formData.zxbm }}
                </el-descriptions-item>
                <el-descriptions-item label="项目状态">
                        {{ formData.xmzt }}
                </el-descriptions-item>
                <el-descriptions-item label="所在城市">
                        {{ formData.xmcs }}
                </el-descriptions-item>
                <el-descriptions-item label="在岸离岸">
                        {{ formData.site }}
                </el-descriptions-item>
                <el-descriptions-item label="结算模式">
                        {{ formData.jsms }}
                </el-descriptions-item>
                <el-descriptions-item label="预算收入">
                        {{ formData.srys }}
                </el-descriptions-item>
                <el-descriptions-item label="预算毛利">
                        {{ formData.mlys }}
                </el-descriptions-item>
                <el-descriptions-item label="预算毛利率">
                        {{ formData.mlyys }}
                </el-descriptions-item>
                <el-descriptions-item label="滚动收入">
                        {{ formData.gdsr }}
                </el-descriptions-item>
                <el-descriptions-item label="滚动毛利">
                        {{ formData.gdml }}
                </el-descriptions-item>
                <el-descriptions-item label="滚动毛利率">
                        {{ formData.gdmll }}
                </el-descriptions-item>
                <el-descriptions-item label="滚动人工">
                        {{ formData.gdrg }}
                </el-descriptions-item>
                <el-descriptions-item label="滚动其他成本">
                        {{ formData.gdqt }}
                </el-descriptions-item>
                <el-descriptions-item label="状态标记">
                        {{ formData.ztbj }}
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
  createMsXmList,
  deleteMsXmList,
  deleteMsXmListByIds,
  updateMsXmList,
  findMsXmList,
  getMsXmListList
} from '@/api/msXmList'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'MsXmList'
})

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               xmbm : [{
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
               xmmc : [{
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
        jssj : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startJssj && !searchInfo.value.endJssj) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startJssj && searchInfo.value.endJssj) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startJssj && searchInfo.value.endJssj && (searchInfo.value.startJssj.getTime() === searchInfo.value.endJssj.getTime() || searchInfo.value.startJssj.getTime() > searchInfo.value.endJssj.getTime())) {
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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
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
  const table = await getMsXmListList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteMsXmListFunc(row)
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
      const res = await deleteMsXmListByIds({ IDs })
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
const updateMsXmListFunc = async(row) => {
    const res = await findMsXmList({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.remsXmList
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMsXmListFunc = async (row) => {
    const res = await deleteMsXmList({ ID: row.ID })
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
  const res = await findMsXmList({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.remsXmList
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
