<template>
    <div class="app-container">
        <p class="page-title">用户列表</p>
        <div class="filter-container">
            <el-button v-permission="`admin-user-add`"
                       type="primary"
                       plain
                       class="filter-item"
                       @click="addGoods()">添加用户
            </el-button>
            <div style="float: right; display: flex">
                <el-input v-model="params.name"
                          placeholder="姓名|手机号"
                          style="width: 200px; margin: 0 3px"
                          class="filter-item" />
                <el-button type="primary"
                           class="filter-item"
                           @click="search()">搜索
                </el-button>
            </div>
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
                             label="用户名"
                             prop="user_name" />
            <el-table-column align="center"
                             label="手机号"
                             prop="tel" />
            <el-table-column align="center"
                             label="真实姓名"
                             prop="real_name" />
            <el-table-column align="center"
                             label="所属角色">
                <template slot-scope="scope">
                    <el-tag v-for="(r, k) in scope.row.roles"
                            :key="k"
                            type="success">{{ r.name }}</el-tag>
                </template>
            </el-table-column>
            <el-table-column align="center"
                             label="状态">
                <template slot-scope="scope">
                    <el-tag v-if="scope.row.status == 1"
                            type="success">正常</el-tag>
                    <el-tag v-else
                            type="danger">停封</el-tag>
                </template>
            </el-table-column>
            <el-table-column align="center"
                             width="170px"
                             label="操作">
                <template slot-scope="scope">
                    <el-button v-permission="`admin-user-edit`"
                               size="mini"
                               type="primary"
                               @click="handleEdit(scope.row)">编辑
                    </el-button>
                    <el-button v-permission="`admin-user-del`"
                               size="mini"
                               type="danger"
                               @click="handleDel(scope.$index, scope.row)">删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
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
import { getList, del } from './api'
import { pageMixin } from '@/utils/mixin'
import checkPermission from '@/utils/permission'
export default {
    name: 'AdminUser',
    mixins: [pageMixin],
    data() {
        return {
            params: {
                name: ''
            }
        }
    },
    methods: {
        checkPermission,
        // 跳转到编辑用户
        handleEdit(info) {
            this.$router.push({ path: '/admin/user/edit', query: { id: info.id }})
        },
        // 跳转到添加用户
        addGoods() {
            this.$router.push('/admin/user/add')
        },
        // 删除用户
        async handleDel(index, info) {
            await this.$confirm('删除用户不可恢复', '警告')
            const { msg } = await del(info.id)
            this.$message.success(msg)
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
        }
    }
}
</script>
