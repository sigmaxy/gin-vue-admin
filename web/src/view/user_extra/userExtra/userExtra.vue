
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">{{t('general.search')}}</el-button>
          <el-button icon="refresh" @click="onReset">{{t('general.reset')}}</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">{{t('general.expand')}}</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>{{t('general.collapse')}}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">{{t('general.add')}}</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">{{t('general.delete')}}</el-button>
            <ExportTemplate  template-id="user_extra_UserExtra" />
            <ExportExcel  template-id="user_extra_UserExtra" filterDeleted/>
            <ImportExcel  template-id="user_extra_UserExtra" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="id" prop="id" width="120" />

            <el-table-column align="left" label="系统用户 ID" prop="sysUserId" width="120" />

            <el-table-column align="left" label="agentCode字段" prop="agentCode" width="120" />

            <el-table-column align="left" label="agentName字段" prop="agentName" width="120" />

            <el-table-column align="left" label="serviceType字段" prop="serviceType" width="120" />

            <el-table-column align="left" label="uccId字段" prop="uccId" width="120" />

        <el-table-column align="left" :label="t('general.operations')" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>{{t('general.detail')}} </el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateUserExtraFunc(scope.row)">{{t('general.edit')}}</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">{{t('general.delete')}}</el-button>
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
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?t('general.add'):t('general.edit')}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">{{t('general.confirm')}}</el-button>
                  <el-button @click="closeDialog">{{t('general.cancel')}}</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="agentCode字段:" prop="agentCode">
    <el-input v-model="formData.agentCode" :clearable="true" placeholder="请输入agentCode字段" />
</el-form-item>
            <el-form-item label="agentName字段:" prop="agentName">
    <el-input v-model="formData.agentName" :clearable="true" placeholder="请输入agentName字段" />
</el-form-item>
            <el-form-item label="serviceType字段:" prop="serviceType">
    <el-input v-model="formData.serviceType" :clearable="true" placeholder="请输入serviceType字段" />
</el-form-item>
            <el-form-item label="uccId字段:" prop="uccId">
    <el-input v-model="formData.uccId" :clearable="true" placeholder="请输入uccId字段" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" :title="t('general.detail')">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="id">
    {{ detailFrom.id }}
</el-descriptions-item>
                    <el-descriptions-item label="系统用户 ID">
    {{ detailFrom.sysUserId }}
</el-descriptions-item>
                    <el-descriptions-item label="agentCode字段">
    {{ detailFrom.agentCode }}
</el-descriptions-item>
                    <el-descriptions-item label="agentName字段">
    {{ detailFrom.agentName }}
</el-descriptions-item>
                    <el-descriptions-item label="serviceType字段">
    {{ detailFrom.serviceType }}
</el-descriptions-item>
                    <el-descriptions-item label="uccId字段">
    {{ detailFrom.uccId }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createUserExtra,
  deleteUserExtra,
  deleteUserExtraByIds,
  updateUserExtra,
  findUserExtra,
  getUserExtraList
} from '@/api/user_extra/userExtra'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'UserExtra'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            agentCode: '',
            agentName: '',
            serviceType: '',
            uccId: '',
        })



// 验证规则
const rule = reactive({
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
  const table = await getUserExtraList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
        confirmButtonText: t('general.confirm'),
         cancelButtonText: t('general.cancel'),
        type: 'warning'
    }).then(() => {
            deleteUserExtraFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: t('general.selectDataToDelete')
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteUserExtraByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.deleteSuccess')
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateUserExtraFunc = async(row) => {
    const res = await findUserExtra({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteUserExtraFunc = async (row) => {
    const res = await deleteUserExtra({ id: row.id })
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

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        agentCode: '',
        agentName: '',
        serviceType: '',
        uccId: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createUserExtra(formData.value)
                  break
                case 'update':
                  res = await updateUserExtra(formData.value)
                  break
                default:
                  res = await createUserExtra(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: t('general.createUpdateSuccess')
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findUserExtra({ id: row.id })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>