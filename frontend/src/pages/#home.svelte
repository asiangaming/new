<script>
    import { initializeApp } from "firebase/app";
    import { getDatabase, ref, onValue } from "firebase/database";
    import dayjs from "dayjs";
    import Carousel from '../lib/Carousel.svelte'

    const firebaseConfig = {
        apiKey: "AIzaSyBqVRbGvJBb1JEfKYyN6jgocZjzsx2lN2A",
        authDomain: "dazzling-pillar-328210.firebaseapp.com",
        databaseURL: "https://dazzling-pillar-328210-default-rtdb.asia-southeast1.firebasedatabase.app",
        projectId: "dazzling-pillar-328210",
        storageBucket: "dazzling-pillar-328210.appspot.com",
        messagingSenderId: "770359422621",
        appId: "1:770359422621:web:7933922e00547dc735ee74"
    };
    const app = initializeApp(firebaseConfig);
    const db = getDatabase(app);
    const sdsb4dday = ref(db, "sdsb4dday");
    const sdsb4dnight = ref(db, "sdsb4dnight");
    
    
    let day_date_draw = "";
    let day_next_year = "";
    let day_next_month = "";
    let day_next_day = "";
    let day_prize1 = "";
    let day_prize2 = "";
    let day_prize3 = "";
    let day_1_prize1 = "0";
    let day_2_prize1 = "0";
    let day_3_prize1 = "0";
    let day_4_prize1 = "0";

    let day_img_1_prize1 = "images/bungapink-blank.png";
    let day_img_2_prize1 = "images/bungapeach-blank.png";
    let day_img_3_prize1 = "images/bungahijau-blank.png";
    let day_img_4_prize1 = "images/bungaungu-blank.png";
    let day_img_1_prize2 = "images/bolakecilbiru-blank.png";
    let day_img_2_prize2 = "images/bolakecilpink-blank.png";
    let day_img_3_prize2 = "images/bolakecilmerah-blank.png";
    let day_img_4_prize2 = "images/bolakecilungu-blank.png";
    let day_img_1_prize3 = "images/bolakecilbiru-blank.png";
    let day_img_2_prize3 = "images/bolakecilpink-blank.png";
    let day_img_3_prize3 = "images/bolakecilmerah-blank.png";
    let day_img_4_prize3 = "images/bolakecilungu-blank.png";
    let night_date_draw = "";
    let night_next_draw = "";
    let night_next_year = "";
    let night_next_month = "";
    let night_next_day = "";
    let night_prize1 = "";
    let night_prize2 = "";
    let night_prize3 = "";
    let night_img_1_prize1 = "number/ball-null.svg";
    let night_img_2_prize1 = "number/ball-null.svg";
    let night_img_3_prize1 = "number/ball-null.svg";
    let night_img_4_prize1 = "number/ball-null.svg";
    let night_img_1_prize2 = "number/ball-null.svg";
    let night_img_2_prize2 = "number/ball-null.svg";
    let night_img_3_prize2 = "number/ball-null.svg";
    let night_img_4_prize2 = "number/ball-null.svg";
    let night_img_1_prize3 = "number/ball-null.svg";
    let night_img_2_prize3 = "number/ball-null.svg";
    let night_img_3_prize3 = "number/ball-null.svg";
    let night_img_4_prize3 = "number/ball-null.svg";
    let temp_day = "";
    let temp_day_hour = "00";
    let temp_day_minute = "00";
    let temp_day_second = "00";
    onValue(sdsb4dday, (snapshot) => {
        const data = snapshot.val();
        day_date_draw = data["datedraw"];
        day_next_year = dayjs(data["nextdraw"]).format("YYYY");
        day_next_month = dayjs(data["nextdraw"]).format("MM");
        day_next_day = dayjs(data["nextdraw"]).format("DD");
        day_prize1 = data["prize1"];
        day_prize2 = data["prize2"];
        day_prize3 = data["prize3"];
        day_1_prize1 = day_prize1[0];
        day_2_prize1 = day_prize1[1];
        day_3_prize1 = day_prize1[2];
        day_4_prize1 = day_prize1[3];
        day_img_1_prize1 = getImage(day_prize1[0],"bungapink");
        day_img_2_prize1 = getImage(day_prize1[1],"bungapeach");
        day_img_3_prize1 = getImage(day_prize1[2],"bungahijau");
        day_img_4_prize1 = getImage(day_prize1[3],"bungaungu");
        day_img_1_prize2 = getImage(day_prize2[0],"bolakecilbiru");
        day_img_2_prize2 = getImage(day_prize2[1],"bolakecilpink");
        day_img_3_prize2 = getImage(day_prize2[2],"bolakecilmerah");
        day_img_4_prize2 = getImage(day_prize2[3],"bolakecilungu");
        day_img_1_prize3 = getImage(day_prize3[0],"bolakecilbiru");
        day_img_2_prize3 = getImage(day_prize3[1],"bolakecilpink");
        day_img_3_prize3 = getImage(day_prize3[2],"bolakecilmerah");
        day_img_4_prize3 = getImage(day_prize3[3],"bolakecilungu");
        temp_day = dayjs(data["nextdraw"] + " 15:00:00").valueOf();
        setInterval(function () {
            let now = new Date().getTime();
            let distance = temp_day - now;
            let hours = Math.floor(
                (distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
            );
            let minutes = Math.floor(
                (distance % (1000 * 60 * 60)) / (1000 * 60)
            );
            let seconds = Math.floor((distance % (1000 * 60)) / 1000);
            temp_day_hour = hours;
            temp_day_minute = minutes;
            temp_day_second = seconds;
            if (hours < 10) {
                temp_day_hour = "0" + hours;
            } else {
                temp_day_hour = hours;
            }
            if (minutes < 10) {
                temp_day_minute = "0" + minutes;
            } else {
                temp_day_minute = minutes;
            }
            if (seconds < 10) {
                temp_day_second = "0" + seconds;
            } else {
                temp_day_second = seconds;
            }
            if (distance < 0) {
                clearInterval();
                temp_day_hour = "00";
                temp_day_minute = "00";
                temp_day_second = "00";
            }
            
        }, 1000);
    });
    let temp_night = "";
    let temp_night_hour = "00";
    let temp_night_minute = "00";
    let temp_night_second = "00";
    let tab_day = "bg-[#74aa63]"
    let tab_night = ""
    let panel_day = true
    let panel_night = false
    onValue(sdsb4dnight, (snapshot) => {
        const data = snapshot.val();
        night_date_draw = data["datedraw"];
        night_next_draw = dayjs(data["nextdraw"]).format("DD-MMM-YYYY");
        night_next_year = dayjs(data["nextdraw"]).format("YYYY");
        night_next_month = dayjs(data["nextdraw"]).format("MM");
        night_next_day = dayjs(data["nextdraw"]).format("DD");
        night_prize1 = data["prize1"];
        night_prize2 = data["prize2"];
        night_prize3 = data["prize3"];
        night_img_1_prize1 = getImage(night_prize1[0],"bungapink");
        night_img_2_prize1 = getImage(night_prize1[1],"bungapeach");
        night_img_3_prize1 = getImage(night_prize1[2],"bungahijau");
        night_img_4_prize1 = getImage(night_prize1[3],"bungaungu");
        night_img_1_prize2 = getImage(night_prize2[0],"bolakecilbiru");
        night_img_2_prize2 = getImage(night_prize2[1],"bolakecilpink");
        night_img_3_prize2 = getImage(night_prize2[2],"bolakecilmerah");
        night_img_4_prize2 = getImage(night_prize2[3],"bolakecilungu");
        night_img_1_prize3 = getImage(night_prize3[0],"bolakecilbiru");
        night_img_2_prize3 = getImage(night_prize3[1],"bolakecilpink");
        night_img_3_prize3 = getImage(night_prize3[2],"bolakecilmerah");
        night_img_4_prize3 = getImage(night_prize3[3],"bolakecilungu");
        temp_night = dayjs(data["nextdraw"] + " 21:00:00").valueOf();
        setInterval(function () {
            let now = new Date().getTime();
            let distance = temp_night - now;
            let hours = Math.floor(
                (distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
            );
            let minutes = Math.floor(
                (distance % (1000 * 60 * 60)) / (1000 * 60)
            );
            let seconds = Math.floor((distance % (1000 * 60)) / 1000);
            if (hours < 10) {
                temp_night_hour = "0" + hours;
            } else {
                temp_night_hour = hours;
            }
            if (minutes < 10) {
                temp_night_minute = "0" + minutes;
            } else {
                temp_night_minute = minutes;
            }
            if (seconds < 10) {
                temp_night_second = "0" + seconds;
            } else {
                temp_night_second = seconds;
            }
            if (distance < 0) {
                clearInterval();
                temp_night_hour = "00";
                temp_night_minute = "00";
                temp_night_second = "00";
            }
        }, 1000);
    });
    function getImage(e,tipe) {
        let urlimg = "";
        switch (e) {
            case "0":
                urlimg = "images/"+tipe+"-0.png";
                break;
            case "1":
                urlimg = "images/"+tipe+"-1.png";
                break;
            case "2":
                urlimg = "images/"+tipe+"-2.png";
                break;
            case "3":
                urlimg = "images/"+tipe+"-3.png";
                break;
            case "4":
                urlimg = "images/"+tipe+"-4.png";
                break;
            case "5":
                urlimg = "images/"+tipe+"-5.png";
                break;
            case "6":
                urlimg = "images/"+tipe+"-6.png";
                break;
            case "7":
                urlimg = "images/"+tipe+"-7.png";
                break;
            case "8":
                urlimg = "images/"+tipe+"-8.png";
                break;
            case "9":
                urlimg = "images/"+tipe+"-9.png";
                break;
        }
        return urlimg;
    }
   
    const handleTab = (e) => {
        switch(e){
            case "day":
                tab_day = "bg-[#74aa63]"
                tab_night = ""
                panel_day = true
                panel_night = false
                break;
            case "night":
                tab_day = ""
                tab_night = "bg-[#74aa63]"
                panel_day = false
                panel_night = true
                break;
        }
    };
  </script>
<section class="hidden lg:flex gap-2 my-5">
    <Carousel />
</section>
<section class="lg:mt-16 mb-5 w-full relative">
    <hr class="w-full bg-[pink] h-[2px] ">
    <h2 class="text-[pink] text-lg lg:text-3xl text-center bg-white absolute -top-4 left-10 z-auto ">다음 그림</h2>
</section>
<ul class="flex justify-center w-full gap-5">
    <li on:click={() => {
        handleTab("day");
    }} class="border-2 border-[#74aa63] {tab_day}  py-2 px-5 rounded-lg cursor-pointer">공주의 날</li>
    <li on:click={() => {
        handleTab("night");
    }} class="border-2 border-[#74aa63] {tab_night} py-2 px-5 rounded-lg cursor-pointer">공주의 밤</li>
</ul>
{#if panel_day}
    <section class="border-solid border-4  border-[#74aa63] mt-5 mb-10 rounded-3xl p-2 lg:p-5">
        <div class="flex flex-col my-2">
            <p class="text-[#74aa63] text-lg lg:text-4xl font-poppins font-extrabold">공주의 날</p>
           
            <div class="flex justify-around w-full gap-2  ">
                <span class="countdown font-mono text-5xl  lg:text-9xl bg-[#74aa63] text-white rounded-full p-5 lg:p-10">
                    <span style="--value:{temp_day_hour};"></span>
                </span>
                <span class="countdown font-mono text-5xl lg:text-9xl bg-[#74aa63] text-white rounded-full p-5 lg:p-10">
                    <span style="--value:{temp_day_minute};"></span>
                </span>
                <span class="countdown font-mono text-5xl lg:text-9xl bg-[#74aa63] text-white rounded-full p-5 lg:p-10">
                    <span style="--value:{temp_day_second};"></span>
                </span>
            </div>
        </div>
    </section>
    <section class="my-5 lg:my-16 w-full relative">
        <hr class="w-full bg-[pink] h-[2px] ">
        <h2 class="text-[pink] text-lg lg:text-3xl text-center bg-white absolute -top-4 left-10 z-auto">오늘의 출력</h2>
    </section>
    <section class="my-5 lg:my-16 w-full relative">
        <img class="w-full z-0" src="images/background_xcferm.jpg" alt="">
        <div class="flex justify-items-center  w-full absolute top-2 left-2 lg:top-10 lg:left-10 ">
            <div class="text-white text-xs lg:text-2xl w-full mx-5">1등상</div>
            <div class="text-white text-xs lg:text-2xl w-full text-right mr-5 lg:mr-20">{day_next_year}년 {day_next_month}월 {day_next_day}일</div>
        </div>
        <div class="hidden lg:flex justify-around  w-full absolute top-28 left-0 ">
            <img class="w-[200px]" src="{day_img_1_prize1}" alt="">
            <img class="w-[200px]" src="{day_img_2_prize1}" alt="">
            <img class="w-[200px]" src="{day_img_3_prize1}" alt="">
            <img class="w-[200px]" src="{day_img_4_prize1}" alt="">
        </div>
        <div class="hidden lg:flex justify-around w-full">
            <div class="flex flex-col absolute bottom-16 left-16 gap-2">
                <div class="text-white text-xs lg:text-2xl w-full">2등상</div>
                <div class="flex justify-start gap-20  w-full  ">
                    <img class="" src="{day_img_1_prize2}" alt="">
                    <img class="" src="{day_img_2_prize2}" alt="">
                    <img class="" src="{day_img_3_prize2}" alt="">
                    <img class="" src="{day_img_4_prize2}" alt="">
                </div>
            </div>
            <div class="flex flex-col absolute bottom-16 right-16 gap-2">
                <div class="text-white text-xs lg:text-2xl w-full">3등상</div>
                <div class="flex justify-start gap-20  w-full  ">
                    <img class="" src="{day_img_1_prize3}" alt="">
                    <img class="" src="{day_img_2_prize3}" alt="">
                    <img class="" src="{day_img_3_prize3}" alt="">
                    <img class="" src="{day_img_4_prize3}" alt="">
                </div>
            </div>
        </div>
        <div class="flex lg:hidden justify-around   w-full absolute top-8 left-0 ">
            <img class="w-[50px]" src="{day_img_1_prize1}" alt="">
            <img class="w-[50px]" src="{day_img_2_prize1}" alt="">
            <img class="w-[50px]" src="{day_img_3_prize1}" alt="">
            <img class="w-[50px]" src="{day_img_4_prize1}" alt="">
        </div>
        <div class="flex flex-col lg:hidden  w-full ">
            <div class="flex absolute bottom-8 text-center gap-1 w-full">
                <div class="text-white text-xs lg:text-2xl w-full">2등상</div>
                <div class="flex justify-start gap-2  w-full  ">
                    <img class="w-[20px]" src="{day_img_1_prize2}" alt="">
                    <img class="w-[20px]" src="{day_img_2_prize2}" alt="">
                    <img class="w-[20px]" src="{day_img_3_prize2}" alt="">
                    <img class="w-[20px]" src="{day_img_4_prize2}" alt="">
                </div>
            </div>
            <div class="flex absolute bottom-2 text-center gap-1 w-full">
                <div class="text-white text-xs lg:text-2xl w-full">3등상</div>
                <div class="flex justify-start gap-2  w-full  ">
                    <img class="w-[20px]" src="{day_img_1_prize3}" alt="">
                    <img class="w-[20px]" src="{day_img_2_prize3}" alt="">
                    <img class="w-[20px]" src="{day_img_3_prize3}" alt="">
                    <img class="w-[20px]" src="{day_img_4_prize3}" alt="">
                </div>
            </div>
        </div>
    </section>
{/if}
{#if panel_night}
    <section class="border-solid border-4  border-[#74aa63] mt-5 mb-10 rounded-3xl  p-5">
        <div class="flex flex-col my-2">
            <p class="text-[#74aa63] text-4xl font-poppins font-extrabold">공주의 밤</p>
            <div class="w-full relative mt-10">
                <img class="w-full" src="images/day.png" alt="">
                <div class="flex justify-center  absolute top-5 w-full gap-36 -left-11 ">
                    <span class="text-[150px] text-white text-center ">{temp_night_hour}</span>
                    <span class="text-[150px] text-white text-center  ">{temp_night_minute}</span>
                    <span class="text-[150px] text-white text-center ">{temp_night_second}</span>
                </div>
            </div>
        </div>
    </section>
    <section class="my-16 w-full relative">
        <hr class="w-full bg-[pink] h-[2px] ">
        <h2 class="text-[pink] text-3xl text-center bg-white absolute -top-4 left-10 z-auto">오늘의 출력</h2>
    </section>
    <section class="my-16 w-full relative">
        <img class="w-full z-0" src="images/background_xcferm.jpg" alt="">
        <div class="flex justify-items-center  w-full absolute top-10 left-10">
            <div class="text-white text-xs lg:text-2xl w-full mx-5">1등상</div>
            <div class="text-white text-xs lg:text-2xl w-full text-right mr-20">{night_next_year}년 {night_next_month}월 {night_next_day}일</div>
        </div>
        <div class="flex justify-around   w-full absolute top-28 left-0 ">
            <img class="w-[200px]" src="{night_img_1_prize1}" alt="">
            <img class="w-[200px]" src="{night_img_2_prize1}" alt="">
            <img class="w-[200px]" src="{night_img_3_prize1}" alt="">
            <img class="w-[200px]" src="{night_img_4_prize1}" alt="">
        </div>
        <div class="flex justify-around w-full">
            <div class="flex flex-col absolute bottom-16 left-16 gap-2">
                <div class="text-white text-xs lg:text-2xl w-full">2등상</div>
                <div class="flex justify-start gap-20  w-full  ">
                    <img class="" src="{night_img_1_prize2}" alt="">
                    <img class="" src="{night_img_2_prize2}" alt="">
                    <img class="" src="{night_img_3_prize2}" alt="">
                    <img class="" src="{night_img_4_prize2}" alt="">
                </div>
            </div>
            <div class="flex flex-col absolute bottom-16 right-16 gap-2">
                <div class="text-white text-xs lg:text-2xl w-full">3등상</div>
                <div class="flex justify-start gap-20  w-full  ">
                    <img class="" src="{night_img_1_prize3}" alt="">
                    <img class="" src="{night_img_2_prize3}" alt="">
                    <img class="" src="{night_img_3_prize3}" alt="">
                    <img class="" src="{night_img_4_prize3}" alt="">
                </div>
            </div>
        </div>
    </section>
{/if}