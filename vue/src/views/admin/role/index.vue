<template>
    <div class="app-container">
        <p class="page-title">角色列表</p>
        <div class="filter-container">
            <el-button type="primary"
                       plain
                       class="filter-item"
                       @click="addRole()">添加角色
            </el-button>
        </div>
        <el-table v-loading="table_loading"
                  element-loading-text="加载中..."
                  border
                  :data="list">
            <el-table-column label="ID"
                             prop="id"
                             align="center"
                             width="60px" />
            <el-table-column align="center"
                             label="角色名"
                             prop="name" />
            <el-table-column align="center"
                             label="所属上级"
                             prop="parent_name" />
            <el-table-column align="center"
                             width="170px"
                             label="操作">
                <template slot-scope="scope">
                    <el-button size="mini"
                               type="primary"
                               @click="handleEdit(scope.row)">编辑
                    </el-button>
                    <el-button size="mini"
                               type="danger"
                               @click="handleDel(scope.$index, scope.row)">删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>

        <el-dialog :visible.sync="dialogTableVisible"
                   :title="dialogTableitile">
            <el-form ref="elForm"
                     :model="formData"
                     :rules="rules"
                     label-width="100px">
                <el-form-item label="所属上级"
                              prop="pid">
                    <el-cascader v-model="formData.pid"
                                 style="width:100%"
                                 clearable
                                 placeholder="无"
                                 :options="roleOptions"
                                 :show-all-levels="false"
                                 :props="roleProps" />
                </el-form-item>
                <el-form-item label="角色名"
                              prop="name">
                    <el-input v-model="formData.name" />
                </el-form-item>
                <el-form-item label="设置权限"
                              prop="auth">
                    <el-tree ref="tree"
                             :data="authOptions"
                             show-checkbox
                             default-expand-all
                             node-key="id"
                             highlight-current
                             :props="authProps" />
                </el-form-item>
            </el-form>
            <div slot="footer">
                <el-button @click="close()">取消</el-button>
                <el-button type="primary"
                           @click="handelConfirm()">确定</el-button>
            </div>
        </el-dialog>

        <div class="page-container">
            <el-pagination background
                           :current-page.sync="params.page"
                           :page-sizes="page_sizes"
                           :page-size="params.pageSize"
                           layout="total, sizes, prev, pager, next, jumper"
                           :total="total"
                           @size-change="handleSizeChange"
                           @current-change="handleCurrentChange" />
        </div>
    </div>
</template>

<script>
import { getList, del, authTree, add, getDetail, edit } from './api'
import { getRoleTree } from '../api'
import { pageMixin } from '@/utils/mixin'
import { removeEmptyChildren } from '@/utils/index'
const formData = {
    id: 0,
    pid: 0,
    name: '',
    auth: []
}

export default {
    name: 'Role',
    mixins: [pageMixin],
    data() {
        return {
            dialogTableVisible: false,
            dialogTableitile: '',
            isAdd: true,
            formData: {
                id: 0,
                pid: 0,
                name: '',
                auth: []
            },
            rules: {
                name: { required: true, trigger: 'blur', message: '请输入角色名' }
            },
            roleOptions: [],
            roleProps: {
                value: 'id',
                label: 'name',
                expandTrigger: 'hover',
                checkStrictly: true,
                emitPath: false
            },
            authOptions: [],
            authProps: {
                label: 'name'
            }
        }
    },
    created() {
        this.init()
    },
    methods: {
        async init() {
            let { data } = await authTree()
            data = removeEmptyChildren(data)
            this.authOptions = data
        },
        // 添加角色
        async addRole() {
            let { data } = await getRoleTree()
            data = removeEmptyChildren(data);
            [this.dialogTableVisible, this.dialogTableitile, this.roleOptions, this.isAdd] = [true, '添加角色', data, true]
            Object.assign(this.formData, formData)
            await this.$nextTick()
            this.$refs.tree.setCheckedKeys([])
        },
        // 编辑角色
        async handleEdit(info) {
            let { data } = await getRoleTree(info.id)
            data = removeEmptyChildren(data)
            const { 'data': role } = await getDetail(info.id);
            [this.dialogTableVisible, this.dialogTableitile, this.formData, this.roleOptions, this.isAdd] = [true, '编辑角色', role, data, false]
            await this.$nextTick()
            this.$refs.tree.setCheckedKeys(role.base_auth)
        },
        // 删除角色
        async handleDel(index, info) {
            await this.$confirm('删除角色不可恢复', '警告')
            const res = await del(info.id)
            this.$message.success(res.msg)
            this.list.splice(index, 1)
        },
        // 获取数据
        async _getData() {
            const res = await getList(this.params)
            this.list = res.data.data
            this.total = res.data.total
            this.pageSize = res.data.per_page
            this.page = res.data.current_page
            this.table_loading = false
        },
        // 关闭弹窗
        close() {
            this.dialogTableVisible = false
        },
        // 提交
        async handelConfirm() {
            await this.$refs.elForm.validate()
            this.formData.auth = this.$refs.tree.getCheckedKeys().concat(this.$refs.tree.getHalfCheckedKeys())
            let res
            if (this.isAdd) {
                res = await add(this.formData)
            } else {
                res = await edit(this.formData.id, this.formData)
            }
            this.$message.success(res.msg)
            this.dialogTableVisible = false
            this._getData()
        }
    }
}
</script>
