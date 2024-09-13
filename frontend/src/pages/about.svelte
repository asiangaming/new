<script>
    let lang = localStorage.getItem("lang");
    let content_text_1 = "";
    if(lang == null || lang == "english"){
        content_text_1 = "ABOUT";
    }else{
        content_text_1 = "VỀ CHÚNG TÔI";
    }
</script>
<section class="my-5 w-full relative">
    <hr class="w-full bg-[pink] h-[2px] ">
    <h2 class="text-[pink] text-lg lg:text-3xl text-center bg-white absolute -top-4 left-10 z-auto">{content_text_1}</h2>
</section>
<div class="flex flex-row mt-10">
    <section class="w-full self-center hidden lg:block">
        <img  src="images/about-me.png" alt="결공주 Lottery Pools / 집">
    </section>
    <section class="p-2 lg:p-10">
        <p class="text-lg lg:text-2xl text-center font-poppins">
            공주 복권 풀은 공주 복권에 관한 결과와 정보를 제공하기 위해 최선을
            다하고 있습니다. 여기에는 일반 플레이어이든 복권 세계에 완전히
            익숙하지 않은 관계없이 게임에 대해 알아야 할 모든 것이 포함됩니다.
            이 사이트는 2006년 복권 시작 이후 모든 공주 복권 풀 티켓을 포함하여
            완전한 결과 세트에 쉽게 액세스할 수 있도록 합니다. 최신 상품은 각
            추첨 직후 거의 매일 표시됩니다. 종합적인 결과 데이터베이스를 통해
            공주 복권 풀은 획득한 숫자에 대한 자세한 통계를 제공할 수 있습니다.
        </p>
    </section>
</div>
