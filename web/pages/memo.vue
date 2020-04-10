<template>
  <div class="container">
    <div>
      <div>
        <h2>Memo List</h2>

        <div style="margin-bottom: 16px;">
          <input v-model="newMemo" type="text">
          <button v-on:click="add">add</button>
        </div>

        <ul>
          <li
            v-for="(memo, i) in memos"
            v-bind:key="memo.ID"
            style="margin-bottom: 8px;"
          >
            <span
              v-if="!memo.editable"
              key="default"
            >
              {{ memo.Text }}
              <button v-on:click="edit(i)">
                edit
              </button>
              <button v-on:click="remove(i)">
                delete
              </button>
            </span>

            <span
              v-else
              key="edit"
            >
              <input
                v-model="memo.newText"
                type="text"
              >
              <button v-on:click="cancel(i)">cancel</button>
              <button v-on:click="update(i)">update</button>
            </span>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      newMemo: '',
      memos: [
        // APIから受け取るデータは、頭文字が大文字
        // { ID: 1, Text: 'Shopping', editable: false, newText: '' },
      ]
    }
  },
  // asyncDataは、ページ遷移前にAPIからデータを取得し、dataに反映します。
  // 反映したいdataのオブジェクトを返り値にします。
  async asyncData ({ $axios, error }) {
    let memos = []

    // APIにGETメソッドでリクエストを送り、memosを取得します。
    try {
      const { data } = await $axios.get('/memo')
      memos = data
    } catch (err) {
      error({
        statusCode: err.response.status,
        message: err.response.statusText
      })
    }

    // APIから取得するデータは、IDとTextだけですので、editableとnewTextを追加します。
    for (let i = 0; i < memos.length; i++) {
      memos[i] = { ...memos[i], editable: false, newText: '' }
    }

    // 取得したmemoosをdataに反映します。
    return { memos }
  },
  methods: {
    async add () {
      if (!this.newMemo) {
        alert('Text is empty')
        return
      }

      let memo
      // APIにPOSTメソッドでリクエストを送り、memoを作成し、そのデータを取得します。
      try {
        const { data } = await this.$axios.post('/memo', { text: this.newMemo })
        // alert(this.newMemo)
        // alert(data)
        memo = { ...data, editable: false, newText: '' }
      } catch (err) {
        alert('Failed to create a new item')
        return
      }

      // 作成したmemoをmemosに追加します。
      this.memos.push(memo)
      // alert(memo)
    },
    edit (i) {
      const newMemo = JSON.parse(JSON.stringify(this.memos[i]))
      newMemo.newText = newMemo.Text
      newMemo.editable = true
      this.memos.splice(i, 1, newMemo)
    },
    async remove (i) {
      // APIにDELETEメソッドでリクエストを送り、memoを削除します。
      try {
        await this.$axios.delete(`/memo/${this.memos[i].ID}`)
      } catch (err) {
        alert('Failed to delete the item')
        return
      }

      this.memos.splice(i, 1)
    },
    cancel (i) {
      const newMemo = JSON.parse(JSON.stringify(this.memos[i]))
      newMemo.editable = false
      this.memos.splice(i, 1, newMemo)
    },
    async update (i) {
      let memo = JSON.parse(JSON.stringify(this.memos[i]))

      // APIにPUTメソッドでリクエストを送り、memoを更新し、そのデータを取得します。
      try {
        const { data } = await this.$axios.put(`/memo/${memo.ID}`, { text: memo.newText })
        memo = { ...data, editable: false, newText: '' }
      } catch (err) {
        alert('Failed to update the item')
        return
      }

      // memoの更新結果をdataに反映します。
      this.memos.splice(i, 1, memo)
    }
  }
}
</script>
