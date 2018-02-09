<template>
  <span class="progress" :style="{width: percent*364 + &quot;px&quot;}"/>
</template>
<script>
export default {
  props: {
    percent: Number
  }
}
</script>
<style>
  .progress{
    display: inline-block;
    width: 364px;
    height: 26px;
    border-radius: 4px;

    background: #4ee;
  }
</style>
