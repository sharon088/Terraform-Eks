security_group_id = "sg-01ce9a534213c5e9b"
public_subnets    = [aws_subnet.public_eu_central_1a.id, aws_subnet.public_eu_central_1b.id]
private_subnets   = [aws_subnet.private_eu_central_1a.id, aws_subnet.private_eu_central_1b.id]
