<template>
  <div class="hello">
    <h1>{{ msg }}</h1>

    <el-table
      :data="userData"
      style="width: 100%">
      <el-table-column
        prop="ID"
        label="ID"
        width="50">
      </el-table-column>
      <el-table-column
        prop="Name"
        label="Name"
        width="180">
      </el-table-column>
      <el-table-column
        prop="Age"
        label="Age"
        width="180">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
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
