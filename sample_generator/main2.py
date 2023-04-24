from datetime import datetime

# 現在の日時を取得
now = datetime.utcnow()

# 指定された形式で日付を文字列に変換
formatted_date = now.strftime('%Y-%m-%dT%H:%M:%SZ')

print(formatted_date)