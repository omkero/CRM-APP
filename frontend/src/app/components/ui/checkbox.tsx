"use client"
import { Check } from 'lucide-react'
import React, {useState} from 'react'

type Props = {
  checked: boolean
  setChecked: React.Dispatch<boolean>
  onClicked: any
}

function Checkbox({checked, setChecked, onClicked}: Props) {
  const [isChecked, setIsChecked] = useState<boolean>(false)
  return (
    <button 
    className='flex items-center justify-center rounded-lg border-1 border-neutral-300 h-6 w-6  cursor-pointer'
    onClick={onClicked}
    >
        {checked && <Check size={17} />}
    </button>
  )
}

export default Checkbox