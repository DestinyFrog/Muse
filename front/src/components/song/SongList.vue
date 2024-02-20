<template>
	<p v-if="loading">Loading ...</p>
	<p v-if="error">{{ error }}</p>

	<main>
		<table v-if="list" class="table-list">
			<thead>
				<tr>
					<th colspan="3">Songs</th>
				</tr>

				<tr>
					<th v-on:click="showForm = !showForm" colspan="3">{{ showForm ? "-" : "+" }}</th>
				</tr>

			</thead>

			<tbody>
				<SongLine v-for="item in list" :key="item.id" :data="item" />
			</tbody>
		</table>

		<FormAddSong v-show="showForm"/>
		<!-- <FormAddSong /> -->
	</main>

</template>

<script>
import FormAddSong from './FormAddSong.vue'
import SongLine from './SongLine.vue'

export default {
	components: {
		FormAddSong,
		SongLine,
	},

	data() {
		return {
			loading: true,
			error: null,
			list: null,

			showForm: false
		}
	},

	created() {
		this.fetchData()
	},

	methods: {
		async fetchData() {
			try {
				const link = `${process.env.VUE_APP_SERVER_URL}/song`
				const {data} = await this.$http.get(link)

				this.list = data
				this.loading = false
			} catch (err) {
				this.error = err.toString()
			}
		}
	}
}
</script>

<style>
main {
	display: flex;
	flex-direction: column;

	/* width: 100; */

	/* padding: 0px 20px; */
	/* margin: 0px 20px; */

	width: 100%;
}

.table-list {
	background-color: transparent;

	max-width: 400px;
	margin: 0px auto;

	padding: 30px 0px;
}

.table-list thead th {
	background-color: aliceblue;
}

.table-list tbody td {
	background-color: aliceblue;
}

.name-row {
	width: 300px;
}
</style>