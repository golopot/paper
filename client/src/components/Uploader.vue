<template>
  <div>
    <h1>Upload</h1>
    <div>
      <form
        class="upload"
        enctype="multipart/form-data"
      >
        <div>
          <div class="upload-box" v-if="!fileSelected">
            <span class="upload-box-image">
              <span class="icon">
                <svg viewBox="0 0 2048 1792" xmlns="http://www.w3.org/2000/svg"><path d="M1344 864q0-14-9-23l-352-352q-9-9-23-9t-23 9l-351 351q-10 12-10 24 0 14 9 23t23 9h224v352q0 13 9.5 22.5t22.5 9.5h192q13 0 22.5-9.5t9.5-22.5v-352h224q13 0 22.5-9.5t9.5-22.5zm640 288q0 159-112.5 271.5t-271.5 112.5h-1088q-185 0-316.5-131.5t-131.5-316.5q0-130 70-240t188-165q-2-30-2-43 0-212 150-362t362-150q156 0 285.5 87t188.5 231q71-62 166-62 106 0 181 75t75 181q0 76-41 138 130 31 213.5 135.5t83.5 238.5z"/>
                </svg>
              </span>
            </span>
            <div>
              上傳檔案
            </div>
            <div style="font-size: 15px;">
              點選 或 拖曳
            </div>
            <input class="upload-button" type="file" name="file" @change="change">
          </div>
        </div>
      </form>

      <div class="upload-progress">
        <Progress :percent="progress" />
      </div>

      <form>
        <div class="expanded" v-if="fileSelected">
          <div><input v-model="title" name="title" placeholder="標題" ></div>
          <div><input v-model="authors" name="authors" placeholder="作者" ></div>
          <div><textarea v-model="abstract" name="abstract" placeholder="摘要（可不填）" /></div>
          <div><button @click="publish" >發表</button></div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Progress from './Progress.vue'
export default {
  components: {
    Progress
  },
  data: () => ({
    fileSelected: false,
    progress: 0,
    title: '',
    authors: '',
    abstract: '',
    fileID: '',
  }),
  methods: {
    change: function(){
      this.fileSelected = true
      const form = document.querySelector('form.upload')
      const body = new FormData(form)
      axios.post('/api/upload', body, {
        onUploadProgress: (ev) => {
          this.progress = ev.loaded / ev.total
          console.log(this.progress)
        }
      })
        .then(r => {
          this.fileID = r.data.fileID
        })
        .catch(console.error)
    },
    publish: function(ev){
      ev.preventDefault()
      console.log(this)
      fetch('/api/paper', {
        method: 'POST',
        body: JSON.stringify({
          title: this.title,
          authors: this.authors,
          fileID: this.fileID,
        })
      })
        .then( r => r.json())
        .then(console.log)
        .catch(console.error)
    }
  },
}

</script>

<style lang="less" scoped>

  .upload-button{
    margin-left: -100px;
    opacity: 0;
    width: 200px;
    height: 170px;
    position: absolute;
    top: 0;
    cursor: pointer;
  }

  .upload-box{
    position: relative;
  }

  .upload-box-image {
    display: inline-block;
    border: 5px dotted #999;
    border-radius: 14px;
    padding: 14px;

    .icon {
      display: inline-block;
      width: 90px;
      opacity: 0.6;
    }
  }

  .expanded {

    input{
      height: 24px;
    }
    textarea{
      height: 80px;
    }

    input,
    textarea{
      width: 360px;
      margin-bottom: 12px;
    }
  }
</style>
