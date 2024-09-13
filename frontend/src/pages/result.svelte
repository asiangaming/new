<script>
      import dayjs from "dayjs";
    export let path_api = "";
    let listgongjunight = [];
    let tab_day = "bg-[#74aa63]";
    let tab_night = "";

    let listgongjuday = [];
    let day = true;

    let lang = localStorage.getItem("lang");
    let content_text_1 = "日にち";
    let content_text_2 = "結果";

    const handleTab = (e) => {
        switch (e) {
            case "day":
                tab_day = "bg-[#74aa63]";
                tab_night = "";
                if (!day) {
                    day = !day;
                }
                if (listgongjuday.length == 0) {
                    initgongjuDAY();
                }
                break;
            case "night":
                tab_day = "";
                tab_night = "bg-[#74aa63]";
                if (day) {
                    day = !day;
                }
                if (listgongjunight.length == 0) {
                    initgongjuNight();
                }
                break;
        }
    };

    async function initgongjuDAY() {
        const resPasar = await fetch(path_api + "api/listshizoukaday", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({}),
        });
        if (!resPasar.ok) {
            const pasarMessage = `An error has occured: ${resPasar.status}`;
            throw new Error(pasarMessage);
        } else {
            const jsonPasar = await resPasar.json();
            if (jsonPasar.status == 200) {
                let record = jsonPasar.record;
                if (record != null) {
                    for (var i = 0; i < record.length; i++) {
                        listgongjuday = [
                            ...listgongjuday,
                            {
                                tgl: dayjs(record[i]["sdsbday_date"]).format(
                                    "DD-MMM-YYYY"
                                ),
                                prize_1: record[i]["sdsbday_prize1"],
                                prize_2: record[i]["sdsbday_prize2"],
                                prize_3: record[i]["sdsbday_prize3"],
                            },
                        ];
                    }
                } else {
                    alert("Error");
                }
            } else {
                alert("Error");
            }
        }
    }
    async function initgongjuNight() {
        const resPasar = await fetch(path_api + "api/listshizoukanight", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({}),
        });
        if (!resPasar.ok) {
            const pasarMessage = `An error has occured: ${resPasar.status}`;
            throw new Error(pasarMessage);
        } else {
            const jsonPasar = await resPasar.json();
            if (jsonPasar.status == 200) {
                let record = jsonPasar.record;
                if (record != null) {
                    for (var i = 0; i < record.length; i++) {
                        listgongjunight = [
                            ...listgongjunight,
                            {
                                tgl: dayjs(record[i]["sdsbnight_date"]).format(
                                    "DD-MMM-YYYY"
                                ),
                                prize_1: record[i]["sdsbnight_prize1"],
                                prize_2: record[i]["sdsbnight_prize2"],
                                prize_3: record[i]["sdsbnight_prize3"],
                            },
                        ];
                    }
                } else {
                    alert("Error");
                }
            } else {
                alert("Error");
            }
        }
    }

    initgongjuDAY();
</script>

<section class="mt-5 mb-5 w-full relative">
    <hr class="w-full bg-[#00989f] h-[2px] " />
    <h2
        class="text-[#7ea8da] text-lg lg:text-3xl text-center bg-white absolute -top-4 left-10 z-auto"
    >
        {content_text_2}
    </h2>
</section>
<ul class="flex justify-center w-full gap-5">
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <li
        on:click={() => {
            handleTab("day");
        }}
        class="border-2 border-[#74aa63] {tab_day}  py-2 px-5 rounded-lg cursor-pointer"
    >
        デイ
    </li>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <li
        on:click={() => {
            handleTab("night");
        }}
        class="border-2 border-[#74aa63] {tab_night} py-2 px-5 rounded-lg cursor-pointer"
    >
        夕方
    </li>
</ul>



{#if day}
    <div class="overflow-x-auto mt-5 mb-10">
        <table class="table-fixed w-full">
            <thead class="sticky top-0">
                <tr>
                    <th
                        class="text-xs lg:text-lg text-center bg-[#ffda92] text-[#4fb47d]"
                        >{content_text_1}</th
                    >
                    <th
                        class="text-xs lg:text-lg text-center bg-[#ffda92] text-[#4fb47d]"
                        >賞品1位</th
                    >
                    <th
                        class="text-xs lg:text-lg text-center bg-[#ffda92] text-[#4fb47d]"
                        >2等賞</th
                    >
                    <th
                        class="text-xs lg:text-lg text-center bg-[#ffda92] text-[#4fb47d]"
                        >3等賞</th
                    >
                </tr>
            </thead>
            <tbody>
                {#each listgongjuday as rec}
                    <tr>
                        <td
                            class="text-xs lg:text-sm text-center bg-transparent border-[1px] border-[#ffda92] text-[#4fb47d]"
                            >{rec.tgl}</td
                        >
                        <td
                            class="text-xs lg:text-sm text-center bg-transparent border-[1px] border-[#ffda92] text-[#4fb47d]"
                            >{rec.prize_1}</td
                        >
                        <td
                            class="text-xs lg:text-sm text-center bg-transparent border-[1px] border-[#ffda92] text-[#4fb47d]"
                            >{rec.prize_2}</td
                        >
                        <td
                            class="text-xs lg:text-sm text-center bg-transparent border-[1px] border-[#ffda92] text-[#4fb47d]"
                            >{rec.prize_3}</td
                        >
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
{:else}
    <div class="overflow-x-auto mt-5 mb-10">
        <table class="table-fixed w-full">
            <thead class="sticky top-0">
                <tr>
                    <th
                        class="text-xs lg:text-lg text-center bg-[#ffda92] text-[#4fb47d]"
                        >{content_text_1}</th
                    >
                    <th
                        class="text-xs lg:text-lg text-center bg-[#ffda92] text-[#4fb47d]"
                        >賞品1位</th
                    >
                    <th
                        class="text-xs lg:text-lg text-center bg-[#ffda92] text-[#4fb47d]"
                        >2等賞</th
                    >
                    <th
                        class="text-xs lg:text-lg text-center bg-[#ffda92] text-[#4fb47d]"
                        >3等賞</th
                    >
                </tr>
            </thead>
            <tbody>
                {#each listgongjunight as rec}
                    <tr>
                        <td
                           class="text-xs lg:text-sm text-center bg-transparent border-[1px] border-[#ffda92] text-[#4fb47d]"
                            >{rec.tgl}</td
                        >
                        <td
                           class="text-xs lg:text-sm text-center bg-transparent border-[1px] border-[#ffda92] text-[#4fb47d]"
                            >{rec.prize_1}</td
                        >
                        <td
                           class="text-xs lg:text-sm text-center bg-transparent border-[1px] border-[#ffda92] text-[#4fb47d]"
                            >{rec.prize_2}</td
                        >
                        <td
                           class="text-xs lg:text-sm text-center bg-transparent border-[1px] border-[#ffda92] text-[#4fb47d]"
                            >{rec.prize_3}</td
                        >
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
{/if}
