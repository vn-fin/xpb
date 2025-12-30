import os
from typing import Optional
import grpc
from xpb import permission_pb2, permission_pb2_grpc


class PermissionClient:
    """gRPC client to verify Firebase Bearer tokens."""

    def __init__(self, server_address: Optional[str] = None):
        if server_address:
            self.server_address = server_address
        else:
            if 'PERMISSION_GRPC_HOST' not in os.environ:
                raise RuntimeError("Environment variable `PERMISSION_GRPC_HOST` not set")
            self.server_address = os.environ['PERMISSION_GRPC_HOST']

        self.channel: Optional[grpc.Channel] = None
        self.stub: Optional[permission_pb2_grpc.PermissionServiceStub] = None
        self._token: Optional[str] = None

    # ----------------- Connection ----------------- #
    def connect(self):
        if not self.channel:
            self.channel = grpc.insecure_channel(self.server_address)
            self.stub = permission_pb2_grpc.PermissionServiceStub(self.channel)

    def close(self):
        if self.channel:
            self.channel.close()
            self.channel = None
            self.stub = None

    def _ensure_stub(self):
        if not self.stub:
            raise RuntimeError("Client not connected. Call connect() first.")

    # ----------------- Authorization ----------------- #
    def set_token(self, token: str):
        """Set the authorization token for subsequent requests."""
        self._token = token

    def _get_metadata(self) -> list:
        """Get metadata with authorization header if token is set."""
        if self._token:
            return [('authorization', self._token)]
        return []

    # ----------------- API Methods ----------------- #
    def get_user_id_from_access_token(self, access_token: str) -> str:
        """Get user_id from access token using RPC."""
        self._ensure_stub()
        request = permission_pb2.CheckAuthRequest(token=access_token)
        try:
            response = self.stub.GetUserIDFromToken(request, metadata=self._get_metadata())
            return response.user_id
        except Exception as e:
            raise ValueError(f"Error: {e}")


# ----------------- Singleton ----------------- #
_permission_client_instance: Optional[PermissionClient] = None


def get_permission_client() -> PermissionClient:
    """Get or create a singleton permission client instance."""
    global _permission_client_instance
    if _permission_client_instance is None:
        client = PermissionClient()
        client.connect()
        _permission_client_instance = client
    return _permission_client_instance


__all__ = ["get_permission_client", "PermissionClient"]
