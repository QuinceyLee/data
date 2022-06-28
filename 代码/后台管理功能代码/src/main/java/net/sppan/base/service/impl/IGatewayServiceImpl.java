package net.sppan.base.service.impl;

import net.sppan.base.dao.IGatewayDao;
import net.sppan.base.dao.IMachineInfoDao;
import net.sppan.base.dao.support.IBaseDao;
import net.sppan.base.entity.GatewayInfo;
import net.sppan.base.entity.MachineInfo;
import net.sppan.base.service.IGatewayService;
import net.sppan.base.service.IMachineInfoService;
import net.sppan.base.service.support.impl.BaseServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class IGatewayServiceImpl extends BaseServiceImpl<GatewayInfo,Integer> implements IGatewayService {

    @Autowired
    private IGatewayDao iGatewayDao;

    @Override
    public IBaseDao<GatewayInfo, Integer> getBaseDao() {
        return iGatewayDao;
    }
}
