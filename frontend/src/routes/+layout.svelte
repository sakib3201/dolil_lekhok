<script lang="ts">
	import '../app.css';

	let { children } = $props();
	let mobileMenuOpen = $state(false);

	function toggleMobileMenu() {
		mobileMenuOpen = !mobileMenuOpen;
	}
</script>

<nav class="bg-white shadow-lg">
	<div class="max-w-7xl mx-auto px-4">
		<div class="flex justify-between h-16">
			<div class="flex items-center">
				<a href="/" class="text-xl font-bold text-indigo-600">
					Dolil Lekhok
				</a>
			</div>
			
			<div class="md:hidden">
				<button 
					type="button"
					class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500"
					onclick={toggleMobileMenu}
				>
					<span class="sr-only">Open main menu</span>
					<!-- Menu open: "hidden", Menu closed: "block" -->
					<svg class="block h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
					</svg>
					<!-- Menu open: "block", Menu closed: "hidden" -->
					<svg class="hidden h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
					</svg>
				</button>
			</div>
			
			<div class="hidden md:flex items-center space-x-4">
				<a href="/about" class="px-3 py-2 text-gray-700 hover:text-indigo-600 transition duration-150">
					About
				</a>
				<a href="/help" class="px-3 py-2 text-gray-700 hover:text-indigo-600 transition duration-150">
					Help
				</a>
				<a href="/login" class="px-3 py-2 text-gray-700 hover:text-indigo-600 transition duration-150">
					Login
				</a>
			</div>
		</div>

		<!-- Mobile menu -->
		{#if mobileMenuOpen}
			<div class="md:hidden">
				<div class="px-2 pt-2 pb-3 space-y-1">
					<a href="/about" class="block px-3 py-2 text-gray-700 hover:text-indigo-600 hover:bg-gray-50 rounded-md text-base font-medium">
						About
					</a>
					<a href="/login" class="block px-3 py-2 text-gray-700 hover:text-indigo-600 hover:bg-gray-50 rounded-md text-base font-medium">
						Login
					</a>
					<a href="/help" class="block px-3 py-2 text-gray-700 hover:text-indigo-600 hover:bg-gray-50 rounded-md text-base font-medium">
						Help
					</a>
				</div>
			</div>
		{/if}
	</div>
</nav>

<div class="min-h-screen">
	{@render children()}
</div>
