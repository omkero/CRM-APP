"use client"
import React from 'react'
import Image from 'next/image'

type Props = {
    LogoTitle: string
}

function Logo({LogoTitle}: Props) {
  return (
    <div className='w-full h-full flex bg-neutral-200 items-center  px-4 py-2 border-r-2 border-r-neutral-300'>
    <div className='flex items-center justify-center text-center'>
        <Image
        src={require("@/app/assets/logo/v_log.jpg")}
        height={60}
        width={60}
        alt='logo'
        />
       <div className='flex items-center gap-3'>
        <h1 className='text-xl font-bold text-gray-800 text-center'>
         {LogoTitle}
        </h1>
        <p className='text-xs text-neutral-500'>0.0.1</p>
       </div>
    </div>
  </div>
  )
}

export default Logo