package net.sppan.base.controller.admin.monitor;

import net.sppan.base.controller.BaseController;
import net.sppan.base.entity.GatewayInfo;
import net.sppan.base.entity.MachineInfo;
import net.sppan.base.service.IGatewayService;
import net.sppan.base.service.IMachineInfoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
@Controller
@RequestMapping("/admin/gatewayInfo")
public class GateWayController extends BaseController{
    @Autowired
    private IGatewayService iGatewayService;

    @RequestMapping("/list")
    @ResponseBody
    public Page<GatewayInfo> list() {
//		SimpleSpecificationBuilder<Resource> builder = new SimpleSpecificationBuilder<Resource>();
//		String searchText = request.getParameter("searchText");
//		if(StringUtils.isNotBlank(searchText)){
//			builder.add("name", Operator.likeAll.name(), searchText);
//		}
        Page<GatewayInfo> page = iGatewayService.findAll(getPageRequest());
        return page;
    }


    @RequestMapping(value = { "/index" })
    public String index() {
        return "admin/byck/index";
    }

    @RequestMapping(value = { "/index2" })
    public String index2() {
        return "admin/sjjk/index";
    }

    @RequestMapping(value = { "/index3" })
    public String index3() {
        return "admin/sjjk/index2";
    }
}
