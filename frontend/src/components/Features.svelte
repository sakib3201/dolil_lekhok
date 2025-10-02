<script>
    import { fade } from 'svelte/transition';
</script>

<section class="features-section" in:fade>
    <h2 class="section-title">শক্তিশালী বৈশিষ্ট্য বাংলাদেশী আইনজীবীদের জন্য</h2>
    <p class="section-subtitle">আপনার আইনি দলিলগুলো ডিজিটাল করতে যা কিছু প্রয়োজন, সবই এখানে</p>
    <div class="features-grid">
        <div class="feature-card">
            <span class="feature-icon">📝</span>
            <h3>অ্যাডভান্সড বাংলা OCR</h3>
            <p>হস্তলিখিত বাংলা আইনি দলিলগুলো ৯৮% নির্ভুলতার সাথে চিনতে পারে, এমনকি কঠিন অবস্থাতেও।</p>
        </div>
        <div class="feature-card">
            <span class="feature-icon">🤖</span>
            <h3>আইনি ফরম্যাটিং</h3>
            <p>স্বয়ংক্রিয়ভাবে সঠিক আইনি দলিল ফরম্যাটিং এবং গঠন প্রয়োগ করে।</p>
        </div>
        <div class="feature-card">
            <span class="feature-icon">📤</span>
            <h3>বহুমুখী এক্সপোর্ট বিকল্প</h3>
            <p>PDF, Word হিসেবে এক্সপোর্ট করুন বা মূল ফরম্যাটিং সহজে ধরে রাখুন।</p>
        </div>
        <div class="feature-card">
            <span class="feature-icon">🔒</span>
            <h3>ব্যাংক-গ্রেড সুরক্ষা</h3>
            <p>এন্ড-টু-এন্ড এনক্রিপশন এবং আপনার দলিলগুলোর নিরাপদ হ্যান্ডলিং নিশ্চিত করে।</p>
        </div>
    </div>
</section>

<style>
    .features-section {
        padding: 80px 20px;
        background-color: #0F172A; /* Darker background for contrast */
        color: #FFFFFF;
        text-align: center;
    }
    .section-title {
        font-size: 2.5rem;
        font-weight: 600;
        margin-bottom: 1rem;
        color: #6366F1; /* Primary color */
        text-shadow: 0 2px 4px rgba(0,0,0,0.1);
    }
    .section-subtitle {
        font-size: 1.125rem;
        font-weight: 400;
        color: #94A3B8; /* Lighter text for better readability */
        margin-bottom: 4rem;
        max-width: 700px;
        margin-left: auto;
        margin-right: auto;
        line-height: 1.6;
    }
    .features-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
        gap: 2rem;
        padding: 0 20px;
        max-width: 1200px;
        margin: 0 auto;
    }
    .feature-card {
        background: linear-gradient(145deg, #1E293B, #0F172A);
        padding: 2rem;
        border: 1px solid #334155;
        border-radius: 12px;
        transition: all 0.3s ease;
        box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
    }
    .feature-card:hover {
        transform: translateY(-8px);
        border-color: #6366F1; /* Primary color */
        box-shadow: 0 10px 25px -5px rgba(99, 102, 241, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
    }
    .feature-icon {
        font-size: 2.5rem;
        margin-bottom: 1.5rem;
        color: #6366F1; /* Primary color */
        display: inline-block;
        transition: transform 0.3s ease;
    }
    
    .feature-card:hover .feature-icon {
        transform: scale(1.1);
        color: #8B5CF6; /* Secondary color on hover */
    }
    .feature-card h3 {
        font-size: 1.375rem;
        font-weight: 600;
        margin-bottom: 1rem;
        color: #FFFFFF;
        position: relative;
        display: inline-block;
    }
    
    .feature-card h3::after {
        content: '';
        position: absolute;
        bottom: -0.5rem;
        left: 0;
        width: 40px;
        height: 3px;
        background: #8B5CF6; /* Secondary color */
        transition: width 0.3s ease;
    }
    
    .feature-card:hover h3::after {
        width: 60px;
    }
    .feature-card p {
        font-size: 1rem;
        line-height: 1.7;
        font-weight: 400;
        color: #94A3B8;
        margin-top: 1rem;
    }
</style>