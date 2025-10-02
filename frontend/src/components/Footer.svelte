<script>
    let email = "";
  
    function subscribe(e) {
        e.preventDefault();
        if (!email) return;
        alert(`Subscribed with ${email}`);
        email = "";
    }

</script>

<!-- Footer -->
<footer class="mt-auto bg-gray-900 text-gray-300 border-t border-gray-800 relative">
  <div class="absolute top-0 left-0 right-0 h-px bg-gradient-to-r from-transparent via-indigo-500 to-transparent opacity-20"></div>
    <div class="max-w-7xl mx-auto px-6 py-12 grid md:grid-cols-3 gap-10">
      <!-- Brand -->
      <div>
        <h2 class="text-2xl font-bold text-white">Dolil Lekhok</h2>
        <p class="mt-3 text-sm text-gray-400">
          হাতে লেখা বাংলা আইনগত নথি সহজে রূপান্তর করুন ডিজিটাল ফরম্যাটে।
        </p>
      </div>

      <!-- Links -->
      <div class="grid grid-cols-2 gap-8">
        <div>
          <h3 class="font-semibold text-white mb-3">প্রোডাক্ট</h3>
          <ul class="space-y-2 text-sm">
            <li><a href="#features" class="hover:text-white">ফিচার</a></li>
            <li><a href="#pricing" class="hover:text-white">প্রাইসিং</a></li>
            <li><a href="#faq" class="hover:text-white">প্রশ্নোত্তর</a></li>
          </ul>
        </div>
        <div>
          <h3 class="font-semibold text-white mb-3">কোম্পানি</h3>
          <ul class="space-y-2 text-sm">
            <li><a href="#about" class="hover:text-white">আমাদের সম্পর্কে</a></li>
            <li><a href="#contact" class="hover:text-white">যোগাযোগ</a></li>
            <li><a href="#blog" class="hover:text-white">ব্লগ</a></li>
          </ul>
        </div>
      </div>

      <!-- Newsletter -->
      <div>
        <h3 class="font-semibold text-white mb-3">সাবস্ক্রাইব করুন</h3>
        <p class="text-sm text-gray-400 mb-3">
          সর্বশেষ আপডেট এবং টিপস পেতে আমাদের নিউজলেটারে যোগ দিন।
        </p>
        <form on:submit={subscribe} class="flex gap-2">
          <input
            type="email"
            bind:value={email}
            placeholder="আপনার ইমেইল"
            class="w-full px-3 py-2 rounded-lg text-gray-900"
            required
          />
          <button
            type="submit"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition"
          >
            পাঠান
          </button>
        </form>
      </div>
    </div>
    <div
      class="border-t border-gray-800 text-center py-6 text-sm text-gray-500"
    >
      © {new Date().getFullYear()} Dolil Lekhok. সর্বস্বত্ব সংরক্ষিত।
    </div>
  </footer>