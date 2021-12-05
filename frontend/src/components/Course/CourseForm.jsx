import React from 'react'
import { useForm } from "react-hook-form";

function CourseForm({ closeForm , onSubmit}) {
    const { register, handleSubmit } = useForm();

    return (
        <div className="fixed z-10 inset-0 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
            <div className="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true" onClick={closeForm}></div>
                <div className="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
                    <div className="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
                        <div className="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                            <h3 className="text-xl leading-6 font-medium text-blue-900 text-center" id="modal-title">
                                Nhập thông tin khoá học mới
                            </h3>
                            <form className="mt-2 space-y-4">
                                <div className="w-full">
                                    <label htmlFor="" className="text-sm font-bold text-black block">Bộ môn*</label>
                                    <select type="text" {...register("department", { required: true })}
                                        className="w-full p-2 border border-gray-300 rounded mt-1" >
                                        <option value="Computer Science">Computer Science</option>
                                        <option value="Software Engineering">Software Engineering</option>
                                    </select>
                                </div>
                                <div>
                                    <label htmlFor="" className="text-sm font-bold text-black block">Mã khoá học*</label>
                                    <input type="text" {...register("code", { required: true })}
                                        className="w-full p-2 border border-gray-300 rounded mt-1" placeholder="Ví dụ: CSC10005" />
                                </div>
                                <div>
                                    <label htmlFor="" className="text-sm font-bold text-black block">Tên khoá học*</label>
                                    <input type="text"
                                        {...register("title", { required: true })}
                                        className="w-full p-2 border border-gray-300 rounded mt-1"
                                        placeholder="Ví dụ: Lập trình hướng đối tượng" />
                                </div>
                                <div>
                                    <label htmlFor="" className="text-sm font-bold text-black block">Mô tả khoá học</label>
                                    <textarea
                                        type="text"
                                        {...register("description")}
                                        className="w-full p-2 border border-gray-300 rounded mt-1"
                                        placeholder="Hãy viết giới thiệu tổng quát về khoá học ..." />
                                </div>
                                <div>
                                    <label htmlFor="" className="text-sm font-bold text-black block">Số tín chỉ*</label>
                                    <input type="number" min="2" {...register("credit", { required: true, min: 2 })}
                                        className="w-full p-2 border border-gray-300 rounded mt-1" />
                                </div>
                                <div>
                                    <label htmlFor="" className="text-sm font-bold text-black block">Học kỳ*</label>
                                    <input type="number" min="1" {...register("semester")}
                                        className="w-full p-2 border border-gray-300 rounded mt-1" />
                                </div>
                                <div>
                                    <label htmlFor="" className="text-sm font-bold text-black block">Năm học</label>
                                    <input type="text"
                                        {...register("academicYear")}
                                        className="w-full p-2 border border-gray-300 rounded mt-1"
                                        placeholder="Ví dụ: 2020-2021" />
                                </div>
                            </form>
                        </div>
                    </div>
                    <div className="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
                        <button type="button"
                            onClick={handleSubmit(onSubmit)}
                            className="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-green-500 text-base font-medium text-white hover:bg-green-7000 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm">
                            Tạo mới
                        </button>
                        <button type="button"
                            onClick={closeForm}
                            className="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm">
                            Cancel
                        </button>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default CourseForm;