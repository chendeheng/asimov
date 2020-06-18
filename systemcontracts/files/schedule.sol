pragma solidity 0.4.25;


interface MarketMaking {
    function exchange(address receiver) external payable returns(uint, uint);
}

contract Schedule {
   address constant delegateAddr = 0x63000000000000000000000000000000000000006d;

   /// initialized or not
    bool private initialized;
    
    function() public payable{}

    function init() public {
        require(!initialized, "it is not allowed to init more than once");

        initialized = true;
    }
    
    function exchange(address marketMakingAddress) public payable {
        require(msg.value > 0, "asset value must bigger than 0");
        
        MarketMaking marketMaking = MarketMaking(marketMakingAddress);
        uint returnValue;
        uint returnAsset;
        (returnValue, returnAsset) = marketMaking.exchange.value(msg.value, msg.assettype)(this);

        msg.sender.transfer(returnValue, returnAsset);
    }
    
}
