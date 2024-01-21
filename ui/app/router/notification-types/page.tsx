"use client"
import { Title } from "@mantine/core";
import { NotificationTypeList } from "@/modules/router/notification-types";
import { useState } from "react";

const NotificationTypesPage = () => {
    const [searchNotificationTypes, setSearchNotificationTypes] = useState('')
    return <>
        <Title order={1}>Notification Types</Title>
        <NotificationTypeList search={searchNotificationTypes} setSearch={setSearchNotificationTypes}/>
    </>
}
export default NotificationTypesPage