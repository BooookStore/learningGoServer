<template>
  <div class="hello">
    <h1>{{ msg }}</h1>

    <h2>User List</h2>
    <el-row>
      <el-col :span="24">
        <el-table
          :data="userData">
          <el-table-column
            prop="ID"
            label="ID">
          </el-table-column>
          <el-table-column
            prop="Name"
            label="Name">
          </el-table-column>
          <el-table-column
            prop="Age"
            label="Age">
          </el-table-column>
        </el-table>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import ElRow from 'element-ui/packages/row/src/row'

export default {
  components: {ElRow},
  name: 'User',
  data () {
    return {
      msg: 'User',
      userData: [{
        ID: 1,
        Name: 'BookStore',
        Age: 24
      }, {
        ID: 2,
        Name: 'BookStore',
        Age: 24
      }]
    }
  }
}
</script>

<style scoped>
  h1, h2 {
    font-weight: normal;
  }
</style>
