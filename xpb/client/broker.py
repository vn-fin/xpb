import os
from typing import Optional
import grpc
from xpb import brokers_pb2, brokers_pb2_grpc

import os
from typing import Optional, List
import grpc
from xpb.brokers import brokers_pb2, brokers_pb2_grpc


class BrokerClient:
    """gRPC client for BrokerGatewayService."""

    def __init__(self, server_address: Optional[str] = None, auth_token: Optional[str] = None):
        """
        Initialize BrokerClient.

        Args:
            server_address: gRPC server address (default: from BROKER_GRPC_HOST env)
            auth_token: Authorization token (default: from BROKER_AUTH_TOKEN env)
        """
        self.server_address = server_address or os.environ.get('BROKER_GRPC_HOST', 'localhost:50051')
        self.auth_token = auth_token or os.environ.get('BROKER_AUTH_TOKEN')
        self.channel: Optional[grpc.Channel] = None
        self.stub: Optional[brokers_pb2_grpc.BrokerGatewayServiceStub] = None

    # ----------------- Connection ----------------- #
    def connect(self):
        """Establish gRPC connection."""
        if not self.channel:
            self.channel = grpc.insecure_channel(self.server_address)
            self.stub = brokers_pb2_grpc.BrokerGatewayServiceStub(self.channel)

    def close(self):
        """Close gRPC connection."""
        if self.channel:
            self.channel.close()
            self.channel = None
            self.stub = None

    def _ensure_stub(self):
        if not self.stub:
            raise RuntimeError("Client not connected. Call connect() first.")

    def _get_metadata(self) -> List[tuple]:
        """Get gRPC metadata with auth header."""
        if self.auth_token:
            return [('authorization', self.auth_token)]
        return []

    # ----------------- API Methods ----------------- #
    def login(self, credential_id: str) -> brokers_pb2.LoginResponse:
        """
        Login with credential ID.

        Args:
            credential_id: The credential ID from database

        Returns:
            LoginResponse with success status and message
        """
        self._ensure_stub()
        request = brokers_pb2.LoginRequest(credential_id=credential_id)
        return self.stub.Login(request, metadata=self._get_metadata())

    def create_order(
        self,
        credential_id: str,
        symbol: str,
        side: str,
        order_type: str,
        quantity: int,
        price: float,
        session_id: str = "",
        user_id: str = ""
    ) -> brokers_pb2.CreateOrderResponse:
        """
        Create a new order.

        Args:
            credential_id: The credential ID
            symbol: Stock symbol (e.g., "VNM")
            side: Order side - "B" (Buy) or "S" (Sell)
            order_type: Order type - "L" (Limit), "M" (Market), "ATO", "ATC"
            quantity: Number of shares
            price: Order price
            session_id: Optional session ID
            user_id: Optional user ID

        Returns:
            CreateOrderResponse with order details
        """
        self._ensure_stub()
        request = brokers_pb2.CreateOrderRequest(
            credential_id=credential_id,
            symbol=symbol,
            side=side,
            type=order_type,
            quantity=quantity,
            price=price,
            session_id=session_id,
            user_id=user_id,
        )
        return self.stub.CreateOrder(request, metadata=self._get_metadata())

    def get_order_by_id(self, credential_id: str, order_id: str) -> brokers_pb2.GetOrderByIdResponse:
        """
        Get order by ID.

        Args:
            credential_id: The credential ID
            order_id: The order ID

        Returns:
            GetOrderByIdResponse with order details
        """
        self._ensure_stub()
        request = brokers_pb2.GetOrderByIdRequest(
            credential_id=credential_id,
            order_id=order_id,
        )
        return self.stub.GetOrderById(request, metadata=self._get_metadata())

    def get_pending_orders(self, credential_id: str) -> brokers_pb2.GetPendingOrdersResponse:
        """
        Get all pending orders.

        Args:
            credential_id: The credential ID

        Returns:
            GetPendingOrdersResponse with list of pending orders
        """
        self._ensure_stub()
        request = brokers_pb2.GetPendingOrdersRequest(credential_id=credential_id)
        return self.stub.GetPendingOrders(request, metadata=self._get_metadata())

    def cancel_order(self, credential_id: str, order_id: str) -> brokers_pb2.CancelOrderResponse:
        """
        Cancel an order.

        Args:
            credential_id: The credential ID
            order_id: The order ID to cancel

        Returns:
            CancelOrderResponse with success status
        """
        self._ensure_stub()
        request = brokers_pb2.CancelOrderRequest(
            credential_id=credential_id,
            order_id=order_id,
        )
        return self.stub.CancelOrder(request, metadata=self._get_metadata())

    def get_portfolio(self, credential_id: str) -> brokers_pb2.GetPortfolioResponse:
        """
        Get portfolio.

        Args:
            credential_id: The credential ID

        Returns:
            GetPortfolioResponse with portfolio items
        """
        self._ensure_stub()
        request = brokers_pb2.GetPortfolioRequest(credential_id=credential_id)
        return self.stub.GetPortfolio(request, metadata=self._get_metadata())

    def get_account_balance(self, credential_id: str) -> brokers_pb2.GetAccountBalanceResponse:
        """
        Get account balance.

        Args:
            credential_id: The credential ID

        Returns:
            GetAccountBalanceResponse with cash and equity
        """
        self._ensure_stub()
        request = brokers_pb2.GetAccountBalanceRequest(credential_id=credential_id)
        return self.stub.GetAccountBalance(request, metadata=self._get_metadata())

    # ----------------- Context Manager ----------------- #
    def __enter__(self):
        self.connect()
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        self.close()


# ----------------- Singleton ----------------- #
_broker_client_instance: Optional[BrokerClient] = None


def get_broker_client(server_address: Optional[str] = None, auth_token: Optional[str] = None) -> BrokerClient:
    """
    Get singleton BrokerClient instance.

    Args:
        server_address: gRPC server address
        auth_token: Authorization token

    Returns:
        Connected BrokerClient instance
    """
    global _broker_client_instance
    if _broker_client_instance is None:
        client = BrokerClient(server_address, auth_token)
        client.connect()
        _broker_client_instance = client
    return _broker_client_instance


__all__ = ["BrokerClient", "get_broker_client"]
