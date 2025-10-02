<script>
    import { fade } from 'svelte/transition';
</script>

<section class="how-it-works-section" in:fade>
    <h2 class="section-title">কীভাবে কাজ করে</h2>
    <p class="section-subtitle">আপনার দলিলগুলো ডিজিটাল করার জন্য সহজ ৩ ধাপ প্রক্রিয়া</p>
    <div class="steps-grid">
        <div class="step-card">
            <span class="step-icon">1️⃣</span>
            <h3>দলিল আপলোড করুন</h3>
            <p>কোনো হস্তলিখিত বাংলা আইনি দলিলের ছবি আপলোড করুন যেকোনো ফরম্যাটে।</p>
        </div>
        <div class="step-card">
            <span class="step-icon">2️⃣</span>
            <h3>AI রূপান্তর</h3>
            <p>আই প্রযুক্তি দিয়ে টেক্সট কনভার্ট করুন ডিজিটাল ফরম্যাটে।</p>
        </div>
        <div class="step-card">
            <span class="step-icon">3️⃣</span>
            <h3>ব্যবহার করুন</h3>
            <p>আপনার পছন্দমতো ডিজিটাল দলিলটি ডাউনলোড করুন এবং ব্যবহার করুন।</p>
        </div>
    </div>
</section>

<style>
    .how-it-works-section {
        padding: 80px 20px;
        background-color: #0F172A;
        color: #FFFFFF;
        text-align: center;
    }
    .section-title {
        font-size: 2.5rem;
        font-weight: 600;
        margin-bottom: 1rem;
        color: #6366F1;
        text-shadow: 0 2px 4px rgba(0,0,0,0.1);
    }
    .section-subtitle {
        font-size: 1.125rem;
        font-weight: 400;
        color: #94A3B8;
        margin-bottom: 4rem;
        max-width: 700px;
        margin-left: auto;
        margin-right: auto;
        line-height: 1.6;
    }
    .steps-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
        gap: 2rem;
        padding: 0 20px;
        max-width: 1200px;
        margin: 0 auto;
        position: relative;
    }
    
    .steps-grid::before {
        content: '';
        position: absolute;
        top: 2.5rem;
        left: 10%;
        right: 10%;
        height: 2px;
        background: linear-gradient(90deg, #6366F1, #8B5CF6, #06B6D4);
        z-index: 1;
    }
    .step-card {
        background: linear-gradient(145deg, #1E293B, #0F172A);
        padding: 2.5rem 2rem;
        border: 1px solid #334155;
        border-radius: 12px;
        transition: all 0.3s ease;
        position: relative;
        z-index: 2;
        box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
    }
    .step-card:hover {
        transform: translateY(-8px);
        border-color: #6366F1;
        box-shadow: 0 10px 25px -5px rgba(99, 102, 241, 0.1);
    }
    .step-icon {
        font-size: 2.5rem;
        margin-bottom: 1.5rem;
        display: inline-block;
        transition: all 0.3s ease;
    }
    
    .step-card:nth-child(1) .step-icon { color: #6366F1; }
    .step-card:nth-child(2) .step-icon { color: #8B5CF6; }
    .step-card:nth-child(3) .step-icon { color: #06B6D4; }
    .step-card h3 {
        font-size: 1.375rem;
        font-weight: 600;
        margin-bottom: 1rem;
        color: #FFFFFF;
        position: relative;
        padding-bottom: 0.75rem;
    }
    
    .step-card h3::after {
        content: '';
        position: absolute;
        bottom: 0;
        left: 0;
        width: 40px;
        height: 3px;
        background: currentColor;
    }
    .step-card p {
        font-size: 1rem;
        line-height: 1.7;
        font-weight: 400;
        color: #94A3B8;
        margin-top: 1rem;
    }
</style>