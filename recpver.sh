# 运行查看网络状态指令
output=$(/opt/cellframe-node/bin/cellframe-node-cli net -net Backbone get status)

# 提取states下的current状态
current_state=$(echo "$output" | grep -A 1 "states:" | grep -oP '(?<=current: ).*')

# 获取当前北京时间并打印
beijing_time=$(TZ="Asia/Shanghai" date +"%Y-%m-%d %H:%M:%S")
echo "当前北京时间: $beijing_time"

# 打印current状态
echo "当前网络状态: $current_state"

# 判断网络状态
case "$current_state" in
    NET_STATE_ONLINE | NET_STATE_SYNC_CHAINS)
        echo "网络在线，无需操作。"
        ;;
    NET_STATE_LINKS_PREPARE | NET_STATE_LINKS_CONNECTING | NET_STATE_LINKS_ESTABLISHED)
    	echo "网络下线,重启网络中"
    	/opt/cellframe-node/bin/cellframe-node-cli net -net Backbone go offline
    
    	# 等待一分钟
	echo "等待一分钟网络重新上线"
    	sleep 60
    
    	# 重新上线
    	/opt/cellframe-node/bin/cellframe-node-cli net -net Backbone go online
    	echo "网络在线"
        ;;
    NET_STATE_OFFLINE)
        echo "当前状态为NET_STATE_OFFLINE，需要重新上线。"
 	echo "网络重新上线中。"
	# 重新上线
    	/opt/cellframe-node/bin/cellframe-node-cli net -net Backbone go online
        ;;
    *)
        echo "未知状态，不执行任何操作。"
        ;;
esac