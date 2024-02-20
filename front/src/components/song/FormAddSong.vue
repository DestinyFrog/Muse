<script>
	export default {
		data() {
			return {
				loading: true,
				error: null,
				list: null,

				form: {
					name: null,
					autor_id: null,
				},
				file: null
			}
		},

		created() {
			this.fetchData()
		},

		methods: {
			async fetchData() {
				try {
					const {data} = await this.$http.get(`${process.env.VUE_APP_SERVER_URL}/autor`)

					this.list = data
					this.loading = false
				} catch (err) {
					this.error = err.toString()
				}
			},

			handleFile() {
				this.file = this.$refs.file.files[0]
			},
			
			async submit() {
				this.autor_id = parseInt( this.form.autor_id )

				try {
					const {data} = await this.$http.post(`${process.env.VUE_APP_SERVER_URL}/song`,
						JSON.stringify(this.form),
						{
							headers: {
								'Content-Type': 'application/json'
							}
						} )

					const id = data.id

					const form = new FormData()
					form.append('file', this.file)
					await this.$http.post(`${process.env.VUE_APP_SERVER_URL}/song/${id}/file`,
						form,
						{
							headers: {
								"enctype":"multipart/form-data",
							}
						})

				} catch (err) {
					this.error = err.toString()
				}
			}
		}
	}
</script>

<template>
	<div class="add-song">
		<h1>Song</h1>

		<p v-show="error">{{ error }}</p>
		<hr/>

		<label>Name:</label>
		<input type="text" v-model="form.name">

		<label>Autor:</label>
		<p v-if="loading">Loading ...</p>
		<select v-if="list" v-model="form.autor_id">
			<option v-for="item in list" :key="item.id" :value="item.id">
				{{ item.name }}
			</option>
		</select>

		<label>File:</label>
		<input type="file" ref="file" accept="audio/mpeg" @change="handleFile($event)">

		<button @click="submit()">Criar</button>
	</div>
</template>

<style>
.add-song {
	display: flex;
	flex-direction: column;
	align-items: center;
	/* grid-template-columns: 1fr 2fr; */

	padding: 10px;
	margin-bottom: 20px;
	border-radius: 5px;

	max-width: 200px;
	/* width: 100%; */
	background-color: antiquewhite;
}

h1 {
	text-align: center;
	margin: 0px 10px;
}

hr {
	width: 80%;
}

label {
	font-size: 12px;
	align-self: start;
	padding-left: 6px;
}

input, select {
	width: 90%;
	margin-bottom: 20px;
	background-color: aliceblue;
	/* border: 1px solid black; */
	/* border-bottom: 1px solid black; */
}

button {
	text-align: center;
}
</style>