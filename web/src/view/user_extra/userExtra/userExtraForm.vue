
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
        <el-form-item>
          <el-button type="primary" @click="save">{{ t('general.save') }}</el-button>
          <el-button type="primary" @click="back">{{ t('general.back') }}</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createUserExtra,
  updateUserExtra,
  findUserExtra
} from '@/api/user_extra/userExtra'

defineOptions({
    name: 'UserExtraForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUserExtra({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
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
               res = await createUserExtra(formData.value)
               break
             case 'update':
               res = await updateUserExtra(formData.value)
               break
             default:
               res = await createUserExtra(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: t('general.createUpdateSuccess')
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
