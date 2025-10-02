<script>
    import { fade } from 'svelte/transition';
</script>

<section class="pricing-section" in:fade>
    <h2 class="section-title">সহজ, স্বচ্ছ মূল্য নির্ধারণ</h2>
    <p class="section-subtitle">আপনার প্রয়োজনের জন্য উপযুক্ত পরিকল্পনা চয়ন করুন</p>
    <div class="pricing-grid">
        <div class="pricing-card">
            <h3>স্টার্টার</h3>
            <p class="price">৳500/মাস</p>
            <ul>
                <li>50 ডকুমেন্ট/মাস</li>
                <li>বেসিক OCR</li>
                <li>ইমেইল সমর্থন</li>
            </ul>
            <button class="get-started">শুরু করুন</button>
        </div>
        <div class="pricing-card popular">
            <h3>প্রফেশনাল <span>সবচেয়ে জনপ্রিয়</span></h3>
            <p class="price">৳1500/মাস</p>
            <ul>
                <li>200 ডকুমেন্ট/মাস</li>
                <li>অ্যাডভান্সড OCR</li>
                <li>আইনি ফরম্যাটিং</li>
                <li>লাইভ সমর্থন</li>
            </ul>
            <button class="get-started">শুরু করুন</button>
        </div>
        <div class="pricing-card">
            <h3>এন্টারপ্রাইজ</h3>
            <p class="price">৳5000/মাস</p>
            <ul>
                <li>অসীমিত ডকুমেন্ট</li>
                <li>প্রিমিয়াম OCR</li>
                <li>24/7 ফোন সমর্থন</li>
                <li>কাস্টম ইন্টিগ্রেশন</li>
            </ul>
            <button class="get-started">বিক্রয়ের সাথে যোগাযোগ</button>
        </div>
    </div>
</section>

<style>
    .pricing-section {
        padding: 80px 20px;
        background-color: #0F172A;
        color: #FFFFFF;
        text-align: center;
        position: relative;
    }
    
    .pricing-section::before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        height: 1px;
        background: linear-gradient(90deg, transparent, rgba(99, 102, 241, 0.5), transparent);
    }
    .section-title {
        font-size: 2.5rem;
        font-weight: 600;
        margin-bottom: 1rem;
        color: #6366F1;
        text-shadow: 0 2px 4px rgba(0,0,0,0.1);
        position: relative;
        display: inline-block;
    }
    
    .section-title::after {
        content: '';
        position: absolute;
        bottom: -10px;
        left: 50%;
        transform: translateX(-50%);
        width: 60px;
        height: 3px;
        background: #8B5CF6;
        border-radius: 2px;
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
    .pricing-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
        gap: 2rem;
        padding: 0 20px;
        max-width: 1200px;
        margin: 0 auto;
    }
    .pricing-card {
        background: linear-gradient(145deg, #1E293B, #0F172A);
        padding: 2.5rem 2rem;
        border: 1px solid #334155;
        border-radius: 12px;
        transition: all 0.3s ease;
        position: relative;
        overflow: hidden;
        box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
    }
    .pricing-card:hover {
        transform: translateY(-8px);
        border-color: #6366F1;
        box-shadow: 0 10px 25px -5px rgba(99, 102, 241, 0.1);
    }
    .pricing-card h3 {
        font-size: 1.5rem;
        font-weight: 600;
        margin-bottom: 1.5rem;
        color: #FFFFFF;
        position: relative;
        display: inline-block;
    }
    
    .pricing-card h3::after {
        content: '';
        position: absolute;
        bottom: -0.75rem;
        left: 0;
        width: 40px;
        height: 3px;
        background: #6366F1;
        transition: width 0.3s ease;
    }
    
    .pricing-card:hover h3::after {
        width: 60px;
    }
    .pricing-card .price {
        font-size: 2.5rem;
        font-weight: 700;
        color: #6366F1;
        margin: 1.5rem 0;
        position: relative;
        display: inline-block;
    }
    
    .pricing-card .price::before {
        content: '৳';
        font-size: 1.5rem;
        font-weight: 500;
        color: #94A3B8;
        margin-right: 0.25rem;
        vertical-align: super;
    }
    
    .pricing-card .price::after {
        content: '/মাস';
        font-size: 1rem;
        font-weight: 400;
        color: #94A3B8;
        margin-left: 0.25rem;
    }
    .pricing-card ul {
        list-style: none;
        padding: 0;
        text-align: left;
        margin: 2rem 0;
    }
    .pricing-card ul li {
        font-size: 1rem;
        font-weight: 400;
        color: #94A3B8;
        margin-bottom: 1rem;
        padding-left: 1.75rem;
        position: relative;
        line-height: 1.6;
    }
    
    .pricing-card ul li::before {
        content: '✓';
        position: absolute;
        left: 0;
        color: #10B981;
        font-weight: bold;
    }
    .pricing-card.popular {
        border: 2px solid #6366F1;
        position: relative;
        transform: scale(1.05);
        z-index: 1;
    }
    
    .pricing-card.popular::before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        height: 4px;
        background: linear-gradient(90deg, #6366F1, #8B5CF6);
    }
    .pricing-card.popular span {
        background: linear-gradient(90deg, #6366F1, #8B5CF6);
        color: #FFFFFF;
        padding: 0.25rem 1rem;
        font-size: 0.75rem;
        font-weight: 600;
        text-transform: uppercase;
        letter-spacing: 1px;
        position: absolute;
        top: -12px;
        left: 50%;
        transform: translateX(-50%);
        border-radius: 20px;
        box-shadow: 0 4px 6px -1px rgba(99, 102, 241, 0.2);
    }
    .get-started {
        background: linear-gradient(135deg, #6366F1, #8B5CF6);
        color: #FFFFFF;
        border: none;
        padding: 0.75rem 2rem;
        border-radius: 6px;
        cursor: pointer;
        font-size: 1rem;
        font-weight: 600;
        width: 100%;
        transition: all 0.3s ease;
        position: relative;
        overflow: hidden;
        z-index: 1;
    }
    
    .get-started::before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        width: 0;
        height: 100%;
        background: linear-gradient(135deg, #8B5CF6, #6366F1);
        transition: all 0.3s ease;
        z-index: -1;
    }
    
    .get-started:hover::before {
        width: 100%;
    }
    
    .pricing-card.popular .get-started {
        background: linear-gradient(135deg, #8B5CF6, #6366F1);
    }
    
    .pricing-card.popular .get-started::before {
        background: linear-gradient(135deg, #6366F1, #8B5CF6);
    }
    .get-started:hover {
        transform: translateY(-2px);
        box-shadow: 0 10px 20px -5px rgba(99, 102, 241, 0.3);
    }
</style>