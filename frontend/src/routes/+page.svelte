<script lang="ts">
	let files: File[] = [];
	
	async function handleUpload() {
		if (files.length === 0) {
			alert('Please select some images to upload');
			return;
		}
		
		// TODO: Implement the actual upload logic here
		console.log('Files to upload:', files);
	}
</script>

<main class="container mx-auto px-4 py-8">
	<h1 class="text-4xl font-bold text-center mb-8">Welcome to Dolil Lekhok</h1>
	
	<div class="max-w-2xl mx-auto">
		<form onsubmit={handleUpload} class="space-y-6">
			<div class="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center">
				<label for="images" class="block text-gray-600">
					<svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
					</svg>
					<p class="mt-4 text-lg">Drag and drop images here, or</p>
					<input
						type="file"
						id="images"
						multiple
						accept="image/*"
						class="hidden"
						bind:files
						onchange={() => {
							if (files.length > 0) {
								alert(`Selected ${files.length} images`);
							}
						}}
					/>
					<button type="button" class="cursor-pointer mt-2 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
						Browse Files
					</button>
				</label>
			</div>
			
			<button
				type="submit"
				class="w-full cursor-pointer flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
			>
				Upload Images
			</button>
		</form>
	</div>
</main>

<style>
	.container {
		max-width: 1200px;
	}
</style>
