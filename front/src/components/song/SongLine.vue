<template>
	<tr>
		<td class="name-line">{{ data.name }}</td>
		<td class="autor-line">{{ data.autor.name }}</td>
		<td v-on:click="deleteData(data.id)">X</td>
	</tr>
</template>

<script>
export default {
	props: [ 'data' ],

	methods: {
		async deleteData(id) {
			if ( !confirm("Certeza que deseja Excluir esse item ?") )
				return

			try {
				await this.$http.delete(`http://localhost:3333/song/${id}`)
			} catch (err) {
				this.error = err.toString()
			}
		}
	}
}
</script>

<style>
	.name-line {
		min-width: 250px;
	}

	.autor-line {
		min-width: 100px;
	}
</style>