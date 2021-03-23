<template>
  <div class="grid bg-gray-200 h-screen">
    <h1 class="text-3xl font-serif text-pink-500 text-center tracking-widest">👋生成易班请假条 for <span
        class="tracking-wider ">CSUSTer</span></h1>
    <div class="grid overflow-scroll">
      <span class="text-gray-500 text-sm font-semibold "> 名字:</span><input
        class="p-3 outline-none text-gray-700 bg-blue-200 block w-full rounded-2xl border-pink-300 shadow-sm focus:border-indigo-300 focus:ring focus:rounded-2xl focus:ring-blue-100 focus:ring-opacity-2"
        type="text" id="name" name="name" value="" placeholder="Your_Names"><br>
      <span class="text-gray-500 text-sm font-semibold "> 学号:</span>
      <input
          class="p-3 outline-none text-gray-700 bg-blue-200 block w-full rounded-2xl border-pink-300 shadow-sm focus:border-indigo-300 focus:ring focus:rounded-2xl focus:ring-blue-100 focus:ring-opacity-2"
          type="text" id="student_id" name="student_id" value="" placeholder="2019693758123"><br>
      <span class="text-gray-500 text-sm font-semibold ">本人联系电话:</span>
      <input
          class="p-3 outline-none text-gray-700 bg-blue-200 block w-full rounded-2xl border-pink-300 shadow-sm focus:border-indigo-300 focus:ring focus:rounded-2xl focus:ring-blue-100 focus:ring-opacity-2"
          type="text" id="tele_phone" name="tele_phone" value="" placeholder="19574839962"><br>
      <span class="text-gray-500 text-sm font-semibold ">外出地点: </span>
      <input
          class="p-3 outline-none text-gray-700 bg-blue-200 block w-full rounded-2xl border-pink-300 shadow-sm focus:border-indigo-300 focus:ring focus:rounded-2xl focus:ring-blue-100 focus:ring-opacity-2"
          type="text" id="location" name="location" value="" placeholder="就在长沙理工大学附近"><br>
      <span
          class="text-gray-500 text-sm font-semibold ">身份证号（随便输18位就好）:</span>
      <input
          class="p-3 outline-none text-gray-700 bg-blue-200 block w-full rounded-2xl border-pink-300 shadow-sm focus:border-indigo-300 focus:ring focus:rounded-2xl focus:ring-blue-100 focus:ring-opacity-2"
          type="text" id="identity_card" name="identity_card" value="" placeholder="591281230812931123"><br>
      <span class="text-gray-500 text-sm font-semibold "> 请假原因:</span>
      <input
          class="p-3 outline-none text-gray-700 bg-blue-200 block w-full rounded-2xl border-pink-300 shadow-sm focus:border-indigo-300 focus:ring focus:rounded-2xl focus:ring-blue-100 focus:ring-opacity-2"
          type="text" id="reason" name="reason" value="" placeholder="就拿个快递"><br>
      <span
          class="text-gray-500 text-sm font-semibold "> 辅导员名字:</span><input
        class="p-3 outline-none text-gray-700 bg-blue-200 block w-full rounded-2xl border-pink-300 shadow-sm focus:border-indigo-300 focus:ring focus:rounded-2xl focus:ring-blue-100 focus:ring-opacity-2"
        type="text" id="coach_name" name="coach_name" value="" placeholder="肖凤"><br>
      <span
          class="text-gray-500 text-sm font-semibold "> 紧急联系人名字: </span><input
        class="p-3 outline-none text-gray-700 bg-blue-200 block w-full rounded-2xl border-pink-300 shadow-sm focus:border-indigo-300 focus:ring focus:rounded-2xl focus:ring-blue-100 focus:ring-opacity-2"
        type="text" id="emergency_contant" name="emergency_contant" value="" placeholder="小淘气"><br>
      <span
          class="text-gray-500 text-sm font-semibold "> 紧急联系人电话: </span><input
        class="p-3 outline-none text-gray-700 bg-blue-200 block w-full rounded-2xl border-pink-300 shadow-sm focus:border-indigo-300 focus:ring focus:rounded-2xl focus:ring-blue-100 focus:ring-opacity-2"
        type="text" id="emergency_tele" name="emergency_tele" value="" placeholder="19671548861"><br>
      <button class="mx-12 font-serif bg-blue-700 hover:bg-blue-900 text-white font-bold py-3 px-4 rounded"
              type="button" onclick="gen()" value="生成假条链接"> 复制生成的请假条网址
      </button>
      <br>
      <button class="mx-12 font-serif bg-blue-700 hover:bg-blue-900 text-white font-bold py-3 px-4 rounded"
              type="button" onclick="openurl()">打开我的请假条
      </button>
    </div>
  </div>
</template>
